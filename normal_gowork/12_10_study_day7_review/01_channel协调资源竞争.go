package main

import "fmt"
import "time"

var ch = make(chan int) //利用管道的阻塞原理控制同步

func printer(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Printf("\n")
}

func persion1() {
	<-ch
	printer("hello")

}

func persion2() {
	printer("world")
	ch <- 666
}

func main() {
	go persion1()
	go persion2() //开辟两个协程
	for {
		//死循环
	}
}
