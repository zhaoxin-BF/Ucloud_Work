package main

import (
	"fmt"
	//"time"
	"runtime"
)

func main() {
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("go")
		}
	}()

	runtime.Gosched() //仅仅只是让出一次时间片 或者一片时间片
	fmt.Println("hello")
}
