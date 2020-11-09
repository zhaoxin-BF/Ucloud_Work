package main

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2"
	//"time"
	//"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name  string
	Phone string
}

func Insert(tm chan<- int) {
	session, err := mgo.Dial("172.18.183.132:27018")
	if err != nil {
		panic(err)
	}
	defer session.Clone()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("student")
	fmt.Println("开始插入！")
	for i := 1; i <= 10000; i++ { //百万

		c.Insert(&Person{"boreas", "666666666"})
		//if err != nil{
		//	fmt.Println("err = ",err)
		//}
		if i%100 == 0 {
			tm <- 1
		}
	}

	fmt.Println("插入完成！")
}

func Manage(tm chan int) {
	ft := time.Now().UnixNano() / 1e6
	st := time.Now().UnixNano() / 1e6
	i := 1
	for {
		<-tm
		fmt.Printf("取一次           ==== len(tm) = %d   =====", len(tm))
		st = time.Now().UnixNano() / 1e6
		fmt.Printf("插入100条数据耗时：%v ms, 共 %v 条数据,  第%v次插入 ;\n", st-ft, i*100, i)
		ft = st
		i++
	}
}

func main() {

	tm := make(chan int, 100000000) //百万缓存
	fmt.Println("1111111111")
	go Manage(tm)
	fmt.Println("2222222222")
	for i := 0; i < 5; i++ {
		go Insert(tm)
	}
	fmt.Println("3333333333")
	for {
		//死循环
	}
}
