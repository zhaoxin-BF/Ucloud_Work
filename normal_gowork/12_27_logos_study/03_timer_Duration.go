/*
//获得时间间隔，time.Now time.Since 类型time.Duration
*/

package main

import (
	"fmt"
	"time"
)

func timeDuration() time.Duration {
	beginTestTime := time.Now()

	time.Sleep(1 * time.Second)

	tD := time.Since(beginTestTime)

	return tD

}

func main() {
	for {
		fmt.Println(timeDuration())
	}
}
