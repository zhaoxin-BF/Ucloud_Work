package main

import (
	"fmt"
	"time"
)

//定义一个打印机，参数为字符串， 按每个字符打印
//打印机属于公共资源

func Printer(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
}

func person1() {
	Printer("hello")
}

func person2() {
	Printer("world")
}

func main() {

	go person1()
	go person2()
	for {
		//死循环
	}
}
