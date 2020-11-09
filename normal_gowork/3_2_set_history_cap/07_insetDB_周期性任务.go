// package main
//
// import (
// 	"fmt"
// 	"time"
// )
//
// func main() {
// 	timer := time.NewTimer(10 * time.Second)
// 	//timer.Reset(1 * time.Second) //重新设置为1s
//
// 	<-timer.C
// 	fmt.Println("子协程可以打印了， 因为定时器的时间到了")
// }
//
// func main01() {
// 	timer := time.NewTimer(3 * time.Second)
//
// 	go func() {
// 		<-timer.C
// 		fmt.Println("子协程可以打印了， 因为定时器的时间到了")
// 	}()
//
// 	timer.Stop() //停止定时器
//
// 	for {
//
// 	}
// }


package main

import (
	"fmt"
	"time"

)

func main() {
	timer := time.NewTicker(10 * time.Second) //循环定时器
	starttime := 1583250900                   //2020-03-03 23:55:00
	for {
		<-timer.C //循环从channel中取数据
		nowtime := time.Now().Unix()
		judgetime := (int(nowtime) - starttime)%86400
		if judgetime < 10{
			fmt.Println(time.Now().Unix())

		}
	}
}