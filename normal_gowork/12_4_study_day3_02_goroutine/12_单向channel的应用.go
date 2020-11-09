package main

import (
	"fmt"
)

func producter(in chan<- int) {
	for i := 0; i < 10; i++ {
		in <- i * i
	}

	close(in)
}

func consumer(out <-chan int) {
	for data := range out {
		fmt.Println("consumer data = ", data)
	}
}

func main() {
	//创建一个双向通道
	ch := make(chan int)

	//生产者， 往channel写内容，写入
	go producter(ch)

	//消费者， 从channel读取内容， 打印
	consumer(ch)

}
