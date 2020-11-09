/*
	1、工作者模式，扇出模式 Fan-Out
	2、多个 goroutine 可以从相同的 channel 中读数据，利用多核并发完成自身的工作，这就是工作者（workers）模式的由来
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

//一个工作者，循环从taskCh 中取出要处理的数据，知道taskCh 关闭后退出执行
func worker(taskCh <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		task, ok := <-taskCh
		if !ok {
			return
		}

		d := time.Duration(task) * time.Millisecond
		time.Sleep(d)
		fmt.Println("Processing task", task)
	}
}

//pool，任务分发器负责开辟协程, workers工人个数， tasks 任务个数
func pool(wg *sync.WaitGroup, workers, tasks int) {

	//创建任务管道
	taskCh := make(chan int)

	//开辟工人协程个数
	for i := 0; i < workers; i++ {
		go worker(taskCh, wg)
	}

	//给任务chanCh添加任务
	for i := 0; i < tasks; i++ {
		taskCh <- i
	}

	close(taskCh)
}

func main() {
	var wg sync.WaitGroup

	//将会有36个工作者协程
	wg.Add(36)

	//36个工人， 50个任务
	go pool(&wg, 36, 50)

	//阻塞等待工人们操作结束
	wg.Wait()

	fmt.Println("程序执行完毕")
}
