package main

import (
	"fmt"
	"time"
)

func worker() string{
	for {
		time.Sleep(10 * time.Second)
		break
	}
	return "程序执行完毕"

}

func main() {
	fmt.Println("开始执行代码")
	str := worker()//阻塞等待返回值

	fmt.Println(str)
}