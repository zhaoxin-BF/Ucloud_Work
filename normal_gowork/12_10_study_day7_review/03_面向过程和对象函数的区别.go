package main

import "fmt"

//实现两数相加
//面向过程

func add01(a, b int) int {
	return a + b
}

//面向对象， 方法：给某个类型绑定一个函数
type long int

func (tmp long) add02(other long) long {
	return tmp + other
}

func main() {
	var result int
	result = add01(1, 1)
	fmt.Println("result = ", result)

	var tmp long = 1

	res := tmp.add02(1)
	fmt.Printf("res = %T\n", res)
	fmt.Println("res = ", res)
}
