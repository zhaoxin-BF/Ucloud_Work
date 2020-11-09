package main

import "fmt"

func testa() {
	fmt.Println("aaaaaaaaaa")
}
func testb() {
	//fmt.Println("bbbbbbbbbb")
	//显示调用pani函数，导致程序中断
	panic("this is a panic test")
}
func testc() {
	fmt.Println("ccccccccccc")
}

func main() {
	testa()
	testb()
	testc()
}
