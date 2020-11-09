package main

import (
	"fmt"
	"regexp"
)

func main() {
	buf := "3.14 567 agsdg 1.23 7. 8.99 1sd1jg1 6.66 7.8 "

	//正则表达式
	reg := regexp.MustCompile(`\d+\.\d+`)
	if reg == nil {
		fmt.Println("MustCompile err")
		return
	}
	//提取关键信息
	result := reg.FindAllString(buf, -1)
	fmt.Println("result = ", result)
}
