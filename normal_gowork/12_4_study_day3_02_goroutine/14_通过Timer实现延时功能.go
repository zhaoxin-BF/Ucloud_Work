package main

import (
	"fmt"
	"time"
)

func main() {
	//延时两秒后打印一句话
	//1.
	time.Sleep(2 * time.Second)
	fmt.Println("时间到")

	//2.
	timer := time.NewTimer(2 * time.Second)
	<-timer.C
	fmt.Println("time is over")

	//3.
	<-time.After(2 * time.Second) //定时两秒，阻塞2s,往channel写内容
	fmt.Println("Time over")
}
