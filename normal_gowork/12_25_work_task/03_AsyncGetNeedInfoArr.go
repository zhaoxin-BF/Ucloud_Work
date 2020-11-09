package main

import (
	"fmt"
	"sync"

	"gopkg.in/mgo.v2"
)

//
func AsyncGetNeedInfoArr(uDiskIDArr []string) []*limax.NeedInfo {

	var waitGroup sync.WaitGroup      //子协程pool控制器，等待其执行完毕后退出
	var needInfoArr []*limax.NeedInfo //返回值指针
	workspaces := 10                  //工作车间数目

	//获得所有的集群数据库
	mDbInfo := limax.GetMDbInfo()

	//创建查询分发器协程
	go pool(&waitGroup, uDiskIDArr, workspaces, mDbInfo, needInfoArr) //needinfoArr 切片指针如何传递？

	waitGroup.Wait()

	//返回查询结果,是一个切片的地址
	return needInfoArr
}

//1、一个工作分发器收集器，根据车间数maxWorkspace的最大限制,和任务数tasks, 当一个车间生产完毕后，再次给其分配task

//wg *sync.WaitGroup                    分发器协程控制器
//uDiskIDArr []string                   需要查询的uDiskID，根据ID在所有集群数据库里面查询uDisk详细信息
//workspaces int                        工作车间的个数
//mDbInfo []limax.MDbInfo                     所有集群的数据库地址信息
//needInfoArr []*limax.NeedInfo         函数内返回值为指针类型，可向下传参，向上传参，不需要数据拷贝？？？？？？？？？？？？？？？？？？？GC会回收内存吗？返回值传值是否更好？
func pool(waitGroup *sync.WaitGroup, uDiskIDArr []string, workspaces int, mDbInfo []limax.MDbInfo, needInfoArr []*limax.NeedInfo) {
	defer waitGroup.Done()

	var wg sync.WaitGroup  //工作车间协程控制器
	var mwg sync.WaitGroup //管理者（收集数据者）协程控制器
	mwg.Add(1)
	//工作车间信息收集缓冲管道,一次性查询1000个uDisk详细信息
	ch := make(chan limax.NeedInfo, len(uDiskIDArr))

	//开辟收集者协程收集打印数据
	go func(mwg *sync.WaitGroup) {
		defer mwg.Done()
		i := 0
		for needInfo := range ch {
			needInfoArr = append(needInfoArr, needInfo...)
			i++
		}

		logs.Info("uDiskIDArr :", len(uDiskIDArr), "needInfoArr :", len(needInfoArr))
		fmt.Printf("一共%d条数据\n", i)

	}(&mwg)

	//非缓冲管道给生产车间分发任务，根据uDiskID分发任务
	uDiskCh := make(chan string)

	//开辟工作车间
	for i := 0; i < workspaces; i++ {
		wg.Add(1)
		go workspace(&wg, taskCh, mDbInfo, ch)
	}

	//给车间分发任务
	for i := 0; i < len(uDiskIDArr); i++ {
		uDiskCh <- uDiskIDArr[i]
	}

	close(uDiskCh)
	//阻塞等待车间生产完成，关闭数据收集ch
	wg.Wait()
	close(ch)

	//阻塞等待数据收集者协程完成数据汇总
	mwg.Wait()
	fmt.Println("查询完毕")
}

//2、一个工作车间，一个工作车间有多个工人，有maxWorker的限制， 和数据库的多少限制，当查完一个库后，再次给其分配库去查询，直至所有任务完毕后退出
//wg *sync.WaitGroup                          共工作车间协程控制器
//uDiskCh chan string                         领取任务的管道
//mDbInfo []limax.MDbInfo                     所有集群的数据库地址信息
//ch chan limax.NeedInfo                      给分发器返回的数据的通道
func workspace(wg *sync.WaitGroup, uDiskCh chan string, mDbInfo []limax.MDbInfo, ch chan limax.NeedInfo) {
	defer wg.Done()

	for {
		//获得任务uDiskID
		uDiskID, ok := <-uDiskCh
		if !ok {
			return
		}

		//非缓冲管道给工人分发任务，根据mDbInfo分发任务
		mDbInfoCh := make(chan limax.MDbInfo)
		quit := make(chan int)

		//开辟工人协程,nums为一个车间工人协程数
		nums := 10
		for i := 0; i < nums; i++ {
			go worker(uDiskID, mDbInfoCh, ch, quit)
		}

		//给工人分发任务，分发数据库地址
		for i := 0; i < len(mDbInfo); i++ {
			select {
			case <-quit:
				break
			}
			mDbInfoCh <- mDbInfo[i]
		}

		close(mDbInfoCh)
	}
}

//3、一个工人，创建好后就一直干活直到结束
//uDiskID string                              //uDiskID, 通过其查询NeedInfo
//mDbInfoCh chan limax.MDbInfo                //集群数据库通道
//ch chan limax.NeedInfo                      //收集查询到的NeedInfo通道
//quit chan int                                //工人协程通信通道（退出通道）
func worker(uDiskID string, mDbInfoCh chan limax.MDbInfo, ch chan limax.NeedInfo, quit chan int) {
	b := bson.M{"extern_id": uDiskID}
	for {
		//获得任务mDbInfo
		mdbInfo, ok := <-mDbInfoCh
		if !ok {
			return
		}

		//GetNeedInfo函数修改方案
		//查到了就返回：NeedInfo,  true
		//没查到就返回：空NeedInfo,  false
		NeedInfo, ok := GetNeedInfo(mdbInfo, b)
		if ok == true {
			ch <- NeedInfo
			quit <- 1
		}
	}
}
