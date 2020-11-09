package main

import "fmt"

func testa() {
	fmt.Println("aaaaaaaaaa")
}
func testb(x int) {
	//设置recover
	defer func() {
		//recover() //不会让错误导致程序奔溃，直接跳过
		//fmt.Println(recover())//打印panic的错误信息
		if err := recover(); err != nil { //产生了panic异常
			fmt.Println(err)
		}
	}() //别忘了(), 调用此匿名函数

	var a [10]int
	a[x] = 111 //数组越界
}
func testc() {
	fmt.Println("ccccccccccc")
}

func main() {
	testa()
	testb(20)
	testc()
}
