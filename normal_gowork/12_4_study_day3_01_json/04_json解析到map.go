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
	//解析为map,且用到断言时，有一个特例是需要用到 万能指针空接口切片类型  []interface{} 实例65行
	m := make(map[string]interface{}, 4) //创建一个map

	err := json.Unmarshal([]byte(jsonBuf), &m) //必须得取地址，第一个参数的意思是
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	fmt.Printf("jsonBuf = %+v\n", m)

	//2.取出单个数据
	var str string
	// str = m["company"]
	// fmt.Println("str =  ", str)//不能运行，没法转换格式

	//3.类型断言
	// for key, value := range m {
	// 	fmt.Printf("%v ==== > %v\n", key, value)
	// 	if key == "company" {
	// 		str = value //不能用，类型不匹配
	// 		fmt.Println("str = ", str)
	// 	}
	// }

	//3.类型断言
	for key, value := range m {
		fmt.Printf("%v ==== > %v\n", key, value)
		switch data := value.(type) {
		case string:
			str = data
			fmt.Printf("map[%s]的值的类型为string， value = %s\n", key, str)
		case bool:
			fmt.Printf("map[%s]的值的类型为bool， value = %v\n", key, data)
		case float64:
			fmt.Printf("map[%s]的值的类型为float64， value = %v\n", key, data)
		case []interface{}: //万能指针空接口
			fmt.Printf("map[%s]的值的类型为string， value = %v\n", key, data)
		}
	}

}
