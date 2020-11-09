/*
	1、工作者模式下的并发查询数据库
*/

package main

import (
	"fmt"
	"sync"

	"gopkg.in/mgo.v2"
)

func worker(waitGroup *sync.WaitGroup, b bson.M, mDbInfoCh chan limax.MDbInfo, ch chan limax.NeedInfo) {
	defer waitGroup.Done()
	
	for {
		mdbInfo, ok := <-mDbInfoCh
		if !ok {
			return
		}

		//GetNeedInfo函数修改方案
		//查到了就返回：NeedInfo,  true
		//没查到就返回：空NeedInfo,  false
		
		NeedInfo, ok := GetNeedInfo(mdbInfo, b)
		if ok == true {
			close(mDbInfoCh)
			ch <- NeedInfo
		}
	}
}


func pool(waitGroup *sync.WaitGroup, b bson.M, workers int, mDbInfo []limax.MDbInfo, ch chan limax.NeedInfo) {
	mDbInfoCh := make(chan limax.MDbInfo)
	//开辟工作协程
	for i := 0; i < workers; i++ {
		go worker(waitGroup, b, mDbInfoCh, ch)
	}

	//分发任务
	for i := 0; i < len(mDbInfo); i++ {
		mDbInfoCh <- mDbInfo[i]
	}
}

///////////////////////////////////////////////////////////////////////hello world
func AsyncGetNeedInfoArr(uDiskID string) limax.NeedInfo {
	var waitGroup sync.WaitGroup
	b := bson.M{"extern_id": uDiskID}
	var needInfo limax.NeedInfo  //返回值
	workers := 10                //工作者协程数目
	waitGroup.Add(workers)

	//获得所有的集群数据库
	mDbInfo := limax.GetMDbInfo()
	
	ch := make(chan limax.NeedInfo)

	//创建查询分发器协程
	go pool(&waitGroup, b, workers, mDbInfo, ch)
	
	//主协程阻塞等待
	needInfo <- ch:
	close(ch)
	
	waitGroup.Wait()
	
	return neneedInfo
}
