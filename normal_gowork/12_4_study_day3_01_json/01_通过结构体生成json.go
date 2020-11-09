package main

import (
	"encoding/json"
	"fmt"
)

/*
{
	"company":"itcast",
	"subjects":[
		"Go",
		"C++",
		"Python",
		"Test"
	],
	"isok":true,
	"price":666.666
}
*/

//成员变量名首字母大写必须大写
// type IT struct {
// 	Company string
// 	Subject []string
// 	Isok    bool
// 	Price   float64
// }

//二次编码
type IT struct {
	Company string   `json:"-"`       //此字段不会输出到屏幕
	Subject []string `json:"subject"` //二次编码
	Isok    bool     `json:",string"` //转换为字符串输出
	Price   float64  `json:",string"` //同上
}

func main() {
	//定义一个结构体变量，同时初始化
	s := IT{"itcast", []string{"Go", "C++", "Python", "Test"}, true, 666.666}

	//编码， 根据内容生成json文本
	buf, err := json.Marshal(s) //非格式化输出：
	//buf, err := json.MarshalIndent(s, "", "  ") //格式化输出
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	fmt.Println("buf = ", string(buf))
}
