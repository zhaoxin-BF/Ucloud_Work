package main

import (
	"fmt"
	"time"
)

var ch = make(chan string)

func main() {

	defer fmt.Println("主协程也结束")

	go func() {
		defer fmt.Println("子协程调用完毕")

		for i := 0; i < 2; i++ {
			fmt.Println("子协程 i= ", i)
			time.Sleep(time.Second)
		}

		ch <- "我是子协程，我结束了"
	}()

	str := <-ch //没有数据， 阻塞
	fmt.Println("str = ", str)

}
