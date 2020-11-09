package main

import (
	"fmt"
	"strings"
)

func main() {
	//"hellogo"中是否包含“hello", 包含返回true, 不包含返回false
	fmt.Println(strings.Contains("hellogo", "hello"))
	fmt.Println(strings.Contains("hellogo", "abc"))

	//Joins组合
	s := []string{"abc", "hello", "mike", "go"}
	buf := strings.Join(s, "@")
	fmt.Println("buf = ", buf)

	//index 的使用
	fmt.Println(strings.Index("abcdhello", "hello")) //返回下标4
	fmt.Println(strings.Index("abcdhello", "go"))    //不包含返回-1

	//Repeat的使用
	buf = strings.Repeat("go", 3)
	fmt.Println("buf = ", buf) //gogogo

	//Split 以指定的分割符拆分
	buf = "hello@abc@mike"
	tmp := strings.Split(buf, "@")
	fmt.Println("buf = ", tmp)

	//Trim去掉两头空格的字符
	buf = strings.Trim("     aaaa     sfd   sdf", " ") //
	fmt.Printf("buf #%s#\n", buf)

	//去掉空格， 把元素放入切片中
	s3 := strings.Fields("  are u ok ?")
	//fmt.Println("s3 = ", s3)
	for i, data := range s3 {
		fmt.Println(i, ", ", data)
	}
}
