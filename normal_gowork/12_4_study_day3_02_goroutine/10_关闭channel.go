package main

import (
	"fmt"
	//"time"
)

//1、通过range遍历
func main() {
	ch := make(chan int, 0)

	fmt.Printf("len(ch) = %d, cap(ch) = %d\n", len(ch), cap(ch))

	//新建协程
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i //写
		}

		close(ch) //关闭管道之后，无法再次写数据，但是可以读取数据
	}()

	for num := range ch { //当关闭后自动跳出循环
		fmt.Println("num =", num)
	}

}

//2、通过for循环遍历，
func main01() {
	ch := make(chan int, 0)

	fmt.Printf("len(ch) = %d, cap(ch) = %d\n", len(ch), cap(ch))

	//新建协程
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i //写
		}

		close(ch) //关闭管道之后，无法再次写数据，但是可以读取数据
	}()

	for { //如果过 ok为true, 说明管道没有关闭
		if num, ok := <-ch; ok == true {
			fmt.Println("num = ", num)
		} else { //管道关闭
			break
		}
	}

}
