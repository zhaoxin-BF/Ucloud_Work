package main

import (
	"fmt"
	// "time"
)

func main() {
	ch := make(chan int, 100000000) //10亿

	fmt.Printf("len(ch) = %d, cap(ch) = %d\n", len(ch), cap(ch))

	//新建协程
	go func() {
		for i := 0; i < 100000000; i++ {
			// fmt.Println("子协程：i= ", i)
			ch <- 1 //写
		}
	}()

	//延时
	// time.Sleep(2 * time.Second)

	for i := 0; i < 100000000; i++ {
		<-ch //读，阻塞
		if len(ch) == 1000 {
			break
		}
	}
	fmt.Printf("len(ch) = %d, cap(ch) = %d\n", len(ch), cap(ch))
	fmt.Println("运行完毕！")
}
