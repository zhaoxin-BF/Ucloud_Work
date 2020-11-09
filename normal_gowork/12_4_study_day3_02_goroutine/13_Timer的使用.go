package main

import (
	"fmt"
	"time"
)

//验证time.NewTimer时间到了, 只会响应一次
func main() {
	timer := time.NewTimer(2 * time.Second) //2s

	for {
		t := <-timer.C //没有数据前会阻塞
		fmt.Println("时间到了", t)
	}

}

func main01() {
	//创建一个定时器，设置时间为2S, 2后， 往timer通道写内容（即当前时间）
	timer := time.NewTimer(2 * time.Second) //2s后，往timer.C写数据，自动写数据
	fmt.Println("当前时间：", time.Now())

	//2s后， 往timer.c写数据， 有数据后， 就可以读取
	t := <-timer.C //没有数据前会阻塞
	fmt.Println("t = ", t)

}
