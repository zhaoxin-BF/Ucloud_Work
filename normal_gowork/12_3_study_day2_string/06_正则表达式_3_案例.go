package main

import (
	"fmt"
	"regexp"
)

func main() {
	buf := `<meta n
	<div>篮球<div>
	<div>足球<div>
	<div>羽毛球<div>>`

	//正则表达式
	reg := regexp.MustCompile(`<div>(?s:(.*?))<div>`)
	if reg == nil {
		fmt.Println("MustCompile err")
		return
	}
	//提取关键信息
	result := reg.FindAllStringSubmatch(buf, -1)
	fmt.Println("result = ", result)

	for _, text := range result {
		fmt.Println("text[0] = ", text[0]) //带<div><div>
		fmt.Println("text[1] = ", text[1]) //不带<div><div>
	}
}
