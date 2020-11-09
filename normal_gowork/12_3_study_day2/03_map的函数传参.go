package main

import "fmt"

func test(m map[int]string) { //map 是引用传参
	m[0] = "Hello World!"
}

func main() {
	m := map[int]string{0: "hello", 1: "hi", 2: "how", 3: "are", 4: "you"}

	fmt.Println("修改前:", m)

	test(m)
	fmt.Println("修改后:", m)
}
