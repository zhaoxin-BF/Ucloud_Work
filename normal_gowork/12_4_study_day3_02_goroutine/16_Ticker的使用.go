package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTicker(1 * time.Second) //循环定时器

	for i := 0; i < 5; i++ {
		<-timer.C //循环从channel中取数据
		if i == 3 {
			timer.Stop()
			break
		}
		fmt.Println("i = ", i)
	}
}
