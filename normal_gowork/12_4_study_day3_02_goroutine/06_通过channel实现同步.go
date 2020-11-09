package main

import (
	"fmt"
	"time"
)

var ch = make(chan int)

//定义一个打印机，参数为字符串， 按每个字符打印
//打印机属于公共资源

func Printer(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Printf("\n")
}

//person1执行完后才person2执行
func person1() {
	Printer("hello")
	ch <- 666 //往管道里面写数据
}

func person2() {
	<-ch //从管道取数据，没有数据就会阻塞
	Printer("world")
}

func main() {

	go person1()
	go person2()
	for {
		//死循环
	}
}
