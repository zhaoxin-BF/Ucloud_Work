package main

import "fmt"

//函数不能重载，只有方法才能重载
func test(n int) {
	fmt.Println("int = ", n)
}

func test(n string) {
	fmt.Println("string = ", n)
}

func main() {
	//n := 10
	s := "string"

	//test(n)
	test(s)
}
