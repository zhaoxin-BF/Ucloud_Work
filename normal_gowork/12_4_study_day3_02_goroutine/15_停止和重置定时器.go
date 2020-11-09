package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(10 * time.Second)
	//timer.Reset(1 * time.Second) //重新设置为1s

	<-timer.C
	fmt.Println("子协程可以打印了， 因为定时器的时间到了")
}

func main01() {
	timer := time.NewTimer(3 * time.Second)

	go func() {
		<-timer.C
		fmt.Println("子协程可以打印了， 因为定时器的时间到了")
	}()

	timer.Stop() //停止定时器

	for {

	}
}
