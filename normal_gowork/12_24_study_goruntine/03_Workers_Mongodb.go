/*
	1、工作者模式下的并发查询数据库
*/

package main

import (
	"fmt"
	"sync"

	"gopkg.in/mgo.v2"
)

type Person struct {
	Name string
	Age  string
}

func worker(wg *sync.WaitGroup, taskCh chan string, ch chan<- Person) {
	defer wg.Done()

	for {
		//获得任务IP 及 Port
		task, ok := <-taskCh
		if !ok {
			return
		}

		//开始连接数据库，进行查询操作
		session, err := mgo.Dial(task)
		if err != nil {
			panic(err)
		}

		session.SetMode(mgo.Monotonic, true)
		c := session.DB("test").C("student")

		stus := make([]Person, 500) //需大于数据库文档数
		err = c.Find(nil).All(&stus)

		fmt.Println("一共查到数据", len(stus))
		for i := 0; i < len(stus); i++ {
			ch <- stus[i]
		}

		//查询结束，关闭数据库连接句柄
		session.Clone()
	}
}

//ipt为任务切片数组， ch 为数据汇总缓冲管道， workers为工作线程的最大个数
func pool(wg *sync.WaitGroup, workers int, ipt []string, ch chan Person) {

	taskCh := make(chan string)
	//开辟工作协程
	for i := 0; i < workers; i++ {
		go worker(wg, taskCh, ch)
	}

	//分发任务
	for i := 0; i < len(ipt); i++ {
		taskCh <- ipt[i]
	}

	close(taskCh)
}

func main() {

	ch := make(chan Person, 100000)
	var wg sync.WaitGroup
	var mwg sync.WaitGroup
	wg.Add(10)
	mwg.Add(1)

	//创建查询任务数据库IP及Port
	ipt := make([]string, 40) //需查询数据库数目
	ipt[0] = "172.18.183.132:27017"
	ipt[1] = "172.18.183.132:27020"
	ipt[2] = "172.18.183.132:27017"
	ipt[3] = "172.18.183.132:27020"
	ipt[4] = "172.18.183.132:27017"
	ipt[5] = "172.18.183.132:27020"
	ipt[6] = "172.18.183.132:27017"
	ipt[7] = "172.18.183.132:27020"
	ipt[8] = "172.18.183.132:27017"
	ipt[9] = "172.18.183.132:27020"
	ipt[10] = "172.18.183.132:27017"
	ipt[11] = "172.18.183.132:27020"
	ipt[12] = "172.18.183.132:27017"
	ipt[13] = "172.18.183.132:27020"
	ipt[14] = "172.18.183.132:27017"
	ipt[15] = "172.18.183.132:27020"
	ipt[16] = "172.18.183.132:27017"
	ipt[17] = "172.18.183.132:27020"
	ipt[18] = "172.18.183.132:27017"
	ipt[19] = "172.18.183.132:27020"
	ipt[20] = "172.18.183.132:27017"
	ipt[21] = "172.18.183.132:27020"
	ipt[22] = "172.18.183.132:27017"
	ipt[23] = "172.18.183.132:27020"
	ipt[24] = "172.18.183.132:27017"
	ipt[25] = "172.18.183.132:27020"
	ipt[26] = "172.18.183.132:27017"
	ipt[27] = "172.18.183.132:27020"
	ipt[28] = "172.18.183.132:27017"
	ipt[29] = "172.18.183.132:27020"
	ipt[30] = "172.18.183.132:27017"
	ipt[31] = "172.18.183.132:27020"
	ipt[32] = "172.18.183.132:27017"
	ipt[33] = "172.18.183.132:27020"
	ipt[34] = "172.18.183.132:27017"
	ipt[35] = "172.18.183.132:27020"
	ipt[36] = "172.18.183.132:27017"
	ipt[37] = "172.18.183.132:27020"
	ipt[38] = "172.18.183.132:27017"
	ipt[39] = "172.18.183.132:27020"

	//开辟收集者协程收集打印数据
	go func(mwg *sync.WaitGroup) {
		defer mwg.Done()
		i := 0
		for per := range ch {
			fmt.Println(per.Name + ":" + per.Age)
			i++
		}
		fmt.Printf("一共%d条数据\n", i)
		fmt.Println("数据输出完毕")
	}(&mwg)

	//分发器开始分发任务
	go pool(&wg, 10, ipt, ch)

	//阻塞等待工人们操作结束
	wg.Wait()
	close(ch)

	mwg.Wait()
	fmt.Println("查询结束")
}
