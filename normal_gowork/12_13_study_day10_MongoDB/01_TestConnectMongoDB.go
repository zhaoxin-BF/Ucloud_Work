package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//type Person struct {
//	Name  string
//	Phone string
//}

func main01() {
	session, err := mgo.Dial("172.18.183.132:27017") //1、自动连接远程MongoDB,IP+Port
	if err != nil {
		panic(err)
	}
	defer session.Close() //2、程序退出后关闭连接

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("student")
	err = c.Insert(&Person{"Ale", "111111"}, &Person{"Cla", "222222222"})
	if err != nil {
		panic(err)
	}
	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		panic(err)
	}
	fmt.Println("Phone:", result.Phone)
}
