package main

import (
	"fmt"
	"time"
)

func main() {
	//创建定时器，设置时间为两秒，2S后往通道里写内容（内容为当前时间）
	timer := time.NewTimer(2 * time.Second)
	fmt.Println("当前时间：", time.Now())

	//2s后写
	t := <-timer.C //这是一个channel
	fmt.Println("t = ", t)
}
