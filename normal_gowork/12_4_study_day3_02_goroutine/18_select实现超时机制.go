package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)    //数据交互channel
	quit := make(chan bool) //结束控制channel

	//新开一个协程
	go func() {
		for {
			select {
			case num := <-ch:
				fmt.Println("num = ", num)
			case <-time.After(1 * time.Second):
				fmt.Println("超时")
				quit <- true
			}
		}
	}()

	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(time.Second)
	}

	<-quit
	fmt.Println("程序结束")

}
