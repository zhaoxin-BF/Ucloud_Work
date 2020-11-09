package main

import "fmt"

func hello(a, b int) {
	fmt.Println("Hello world!")
	fmt.Printf("a = %d, b = %d\n", a, b)
}
func eachfor(args ...int) {
	for _, data := range args {
		fmt.Printf("data = %d\n", data)
	}
}

func temp(args ...int) {
	eachfor(args...)
}
func main() {
	temp(1, 2, 3, 4, 5, 6)
	hello(100, 200)
	fmt.Println("主函数执行完毕！")

	var (
		a int
		b int
	)

	a, b = 100, 200
	fmt.Printf("a = a%d, b = b%d\n", a, b)

	c, d := 300, 400
	fmt.Printf("c = %d, d = %d", c, d)

	type (
		char  byte
		float int32
	)

	var (
		ch char
		fa float
	)

	ch = 'c'
	fa = 3

	fmt.Printf("ch = %T, fa = %T\n", ch, fa)
}
