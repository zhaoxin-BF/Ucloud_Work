package main

//访问配置文件（conf .json文件）

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"github.com/astaxie/beego/logs"
)

func Check(remoteUser string) string {
	useByte, err := ioutil.ReadFile("./conf/user.json")
	var userMap map[string][]string

	if err != nil{
		return "visitor"
	} else {
		json.Unmarshal(useByte, &userMap)
	}

	fmt.Println(userMap)

	for key, list := range userMap{
		for _, username := range list{
			fmt.Println(username)
			if username == remoteUser{
				fmt.Println(username)
				return key
			}
		}
	}

	return "visitor"
}

func main(){
	identity := Check("han.jia")
	logs.Info("查询到：",identity)
}