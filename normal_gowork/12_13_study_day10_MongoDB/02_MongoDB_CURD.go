package main

import (
	"fmt"
	//"gopkg.in/mgo.v2/bson"

	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name string
	Age  []int
}

////一、插入数据库
//func main() {
//	fmt.Print("开始插入")
//	session, err := mgo.Dial("117.50.104.98:27017")
//	if err != nil {
//		panic(err)
//	}
//	defer session.Clone()
//	var age []int
//	age = append(age, 1)
//	age = append(age, 2)
//	session.SetMode(mgo.Monotonic, true)
//	c := session.DB("test").C("student")
//	for i := 0; i < 10; i++ {
//		err = c.Insert(&Person{"beifeng", age})
//		if err != nil {
//			panic(err)
//		}
//	}
//
//	fmt.Println("插入完毕！")
//}


//一、插入数据库
func main() {
	fmt.Print("开始查找数据库")
	session, err := mgo.Dial("117.50.104.98:27017")
	if err != nil {
		panic(err)
	}
	defer session.Clone()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("student")
	var person []Person

	c.Find(nil).All(&person)
	fmt.Println(person)
	for _, val := range person{
		fmt.Println(val)
	}
}









//二、查找数据库
//func main() {
//	str := "172.18.183.132:27020"
//	session, err := mgo.Dial(str)
//	if err != nil {
//		panic(err)
//	}
//	defer session.Clone()
//
//	session.SetMode(mgo.Monotonic, true)
//	c := session.DB("test").C("student")
//
//	//stus := make([]Person, 200)
//	var stus []Person
//	err = c.Find(nil).All(&stus)
//
//	fmt.Println(stus[1].Name + ":" + stus[1].Age)
//
//	fmt.Println(len(stus), "条数据查找完毕！")
//}

//三、bson查找数据库
//func main() {
//	str := "172.18.183.132:27020"
//	session, err := mgo.Dial(str)
//	if err != nil {
//		panic(err)
//	}
//	defer session.Clone()
//
//	session.SetMode(mgo.Monotonic, true)
//	c := session.DB("test").C("student")
//
//	//stus := make([]Person, 200)
//	var stus []Person
//	fmt.Println(len(stus), "开始查找！")
//	//$group: { _id : null, sum : { $sum: "$length" } }
//	//err = c.Find(nil).Select(bson.M{"age":bson.M{"$sum":"age"}}).All(&summ)
//	err = c.Find(bson.M{"name":"beifeng"}).Select(bson.M{"age":1}).All(&stus)
//
//	fmt.Println(stus)
//
//	fmt.Println(len(stus), "条数据查找完毕！")
//}
