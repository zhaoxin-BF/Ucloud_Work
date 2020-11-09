package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	ch := make(chan string, 10000)

	for i := 0; i < 10; i++ {
		go func() {

			for j := 1; j <= 10; j++ {
				ch <- "子协程写入"
			}
			fmt.Println("子协程写入完成")
			wg.Done()
		}()
	}

	// defer func() {
	// 	for x := range ch {
	// 		fmt.Println(x)
	// 	}

	// 	close(ch)
	// 	fmt.Println("主协程退出")
	// }()

	wg.Wait()
	defer func() {
		n := len(ch)
		for i := 0; i < n; i++ {
			fmt.Printf("第%d条数据：%s\n", i, <-ch)
		}
		close(ch)
		fmt.Println("主协程退出")
	}()

}
