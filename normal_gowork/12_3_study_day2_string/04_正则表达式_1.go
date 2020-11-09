package main

import (
	"fmt"
	"regexp"
)

func main() {
	buf := "abc azc a7c aac 888 a9c tac"

	//(1) 解释规则,这个函数会解析正则表达式，如果成功返回
	reg1 := regexp.MustCompile("a[0-9]c")

	if reg1 == nil { //解析失败，返回nil
		fmt.Println("regexp err")
		return
	}

	//2)根据规则提取关键信息,-1表示全部符合都给出来，1 表示只需要一个
	result := reg1.FindAllStringSubmatch(buf, -1)
	fmt.Println("result1", result)
}
