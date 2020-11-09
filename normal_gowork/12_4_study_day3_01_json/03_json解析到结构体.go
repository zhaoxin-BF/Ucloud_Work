package main

import (
	"encoding/json"
	"fmt"
)

type IT struct { //结构体变量命令必须是大写的
	Company  string   `json:"company"`  //此字段不会输出到屏幕
	Subjects []string `json:"subjects"` //二次编码
	Isok     bool     `json:"isok"`     //转换为字符串输出
	Price    float64  `json:"price"`    //同上
}

func main() {
	//1、json解析到结构体
	jsonBuf := `
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
	}`

	// var tmp IT                                   //定义一个结构体变量
	// err := json.Unmarshal([]byte(jsonBuf), &tmp) //必须得取地址，第一个参数的意思是
	// if err != nil {
	// 	fmt.Println("err = ", err)
	// 	return
	// }

	// fmt.Println("jsonBuf = ", tmp)

	//2、只需要一个json参数
	type IT2 struct {
		Company string `json:"company"` //此字段不会输出到屏幕
	}

	var tmp2 IT2
	err := json.Unmarshal([]byte(jsonBuf), &tmp2) //第一参数为将原生的json数据转化为字节数据传入
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	fmt.Printf("tmp2 = %+v\n", tmp2) //%v自动匹配格式输出
}
