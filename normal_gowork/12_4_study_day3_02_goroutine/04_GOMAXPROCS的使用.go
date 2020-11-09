package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(8) //参数 1 指定以单核运算
	n := runtime.NumCPU()
	fmt.Println("n = ", n)
	for {
		go fmt.Print(1)

		fmt.Print(0)
	}
}
