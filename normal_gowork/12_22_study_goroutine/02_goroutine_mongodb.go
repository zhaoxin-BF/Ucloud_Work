package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name string
	Age  string
}

func main() {
	ch := make(chan Person, 1000)
	str := "172.18.183.132:27017"
	session, err := mgo.Dial(str)
	if err != nil {
		panic(err)
	}
	defer session.Clone()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("student")

	stus := make([]Person, 1000)
	err = c.Find(nil).All(&stus)

	fmt.Println("一共查到数据", len(stus))
	for i := 0; i < len(stus); i++ {
		ch <- stus[i]
	}
	i := 0
	n := len(ch)
	for ; i < n; i++ {
		per := <-ch
		fmt.Println(per.Name + ":" + per.Age)
	}

	fmt.Printf("一共%d条数据", i)
	fmt.Println("查找完毕！")
}
