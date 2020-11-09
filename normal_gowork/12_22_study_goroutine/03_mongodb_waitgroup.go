package main

import (
	"fmt"
	"sync"

	"gopkg.in/mgo.v2"
)

type Person struct {
	Name string
	Age  string
}

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	ch := make(chan Person, 10000) //总文档数预测

	for i := 0; i < 5; i++ { //循环开辟协程
		go func() {
			str := "172.18.183.132:27020"
			session, err := mgo.Dial(str)
			if err != nil {
				panic(err)
			}
			defer session.Clone()

			session.SetMode(mgo.Monotonic, true)
			c := session.DB("test").C("student")

			stus := make([]Person, 1000) //大于文档数
			err = c.Find(nil).All(&stus)

			fmt.Println("一共查到数据", len(stus))
			for i := 0; i < len(stus); i++ {
				ch <- stus[i]
			}
			wg.Done()
		}()
	}

	wg.Wait() //等待子协程运行完毕后退出主协程

	defer func() {
		i := 0
		n := len(ch)
		for ; i < n; i++ {
			per := <-ch
			fmt.Println(per.Name + ":" + per.Age)
		}

		fmt.Printf("一共%d条数据", i)
		fmt.Println("查找完毕！")
		close(ch)
		fmt.Println("主协程退出")
	}()
}
