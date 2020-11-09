/*
	1、扇入模式 Fan-In
	2、简单来说，扇入模式就是一个函数从多个输入源读取数据并且复用到单个 channel 中
*/

package main

import (
	"fmt"
	"time"
)

func producer(ch chan int, d time.Duration) {
	var i int
	for {
		ch <- i
		i++
		time.Sleep(d)
	}
}

func reader(out chan int) {
	for x := range out {
		fmt.Println(x)
	}
}

func main() {
	ch := make(chan int)
	out := make(chan int)

	go producer(ch, 100*time.Millisecond)
	go producer(ch, 250*time.Millisecond)

	go reader(out)

	for i := range ch {
		out <- i
	}
}
