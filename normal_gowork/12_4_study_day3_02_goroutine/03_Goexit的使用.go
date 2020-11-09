package main

import (
	"fmt"
	"runtime"
)

func test() {
	defer fmt.Println("cccccccccccc") //终止函数即调用，终止协程就不再给分配时间片了，就不再执行任何操作

	//return 				//终止次函数
	runtime.Goexit() //终止所在的协程
	fmt.Println("dddddddddddd")
}

func main() {
	go func() {
		fmt.Println("aaaaaaaaaaaa")

		//调用了别的函数
		test()

		fmt.Println("bbbbbbbbbbbbb") //由于子协程终止了， 所以后面就不打印了
	}()

	//主协程死循环
	for i := 0; i > -1; i++ {
	}
}
