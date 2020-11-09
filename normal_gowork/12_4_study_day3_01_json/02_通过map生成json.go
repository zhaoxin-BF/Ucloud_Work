package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//创建一个map
	m := make(map[string]interface{}, 4) //interface{}因为时不同类型，所以要用空接口
	m["company"] = "itcast"
	m["subject"] = []string{"Go", "C++", "Python", "Test"}
	m["isok"] = true
	m["price"] = 666.666

	//result, err := json.Marshal(m)
	result, err := json.MarshalIndent(m, "", "	") //第三个是一个tab键
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	fmt.Println("result = ", string(result))
}
