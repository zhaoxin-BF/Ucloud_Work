package main

import (
	"fmt"
	"strconv"
)

func main() {
	//转换为字符串后追加到数组
	slice := make([]byte, 0, 1024)
	slice = strconv.AppendBool(slice, true)

	//第二个数为要追加的数， 第三个为指定10进制方式追加
	slice = strconv.AppendInt(slice, 1234, 10)
	slice = strconv.AppendQuote(slice, "abcgohello")

	fmt.Println("slie = ", string(slice)) //转换string后在打印

	//其他类型转换为字符串
	var str string
	str = strconv.FormatBool(false)
	//f指打印格式，以小数的格式，-1指小数点位数(紧缩模式)，64以float64处理
	str = strconv.FormatFloat(3.14, 'f', -1, 64)

	//整型转字符串， 常用
	str = strconv.Itoa(6666)
	fmt.Println("str", str)

	//字符串转其他类型
	flag, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println("flag: ", flag)
	} else {
		fmt.Println("err = ", err)
	}

	//把字符串，转换为整型，Atoi
	a, _ := strconv.Atoi("2344")
	fmt.Println("a = ", a)
}
