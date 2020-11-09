/*
	1、利用缓冲channel  ch  控制协程并发量
	2、管理者协程进行汇总数据的打印工作，同步写入输出
	3、新建退出通道，通知管理者协程的退出，结合数据通道ch的长度，进行退出
	4、主协程 同一个waitGroup 进行阻塞等待退出
*/
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
	// var wg sync.WaitGroup
	wg := &sync.WaitGroup{} //值传递，所以传给协程时注意传地址

	quit := make(chan int)         //控制管理者协程退出
	ch := make(chan Person, 10120) //总文档数预测
	goch := make(chan int, 10)     //限制协程最大数目
	ipt := make([]string, 40)      //需查询数据库数目
	ipt[0] = "172.18.183.132:27017"
	ipt[1] = "172.18.183.132:27020"
	ipt[2] = "172.18.183.132:27017"
	ipt[3] = "172.18.183.132:27020"
	ipt[4] = "172.18.183.132:27017"
	ipt[5] = "172.18.183.132:27020"
	ipt[6] = "172.18.183.132:27017"
	ipt[7] = "172.18.183.132:27020"
	ipt[8] = "172.18.183.132:27017"
	ipt[9] = "172.18.183.132:27020"
	ipt[10] = "172.18.183.132:27017"
	ipt[11] = "172.18.183.132:27020"
	ipt[12] = "172.18.183.132:27017"
	ipt[13] = "172.18.183.132:27020"
	ipt[14] = "172.18.183.132:27017"
	ipt[15] = "172.18.183.132:27020"
	ipt[16] = "172.18.183.132:27017"
	ipt[17] = "172.18.183.132:27020"
	ipt[18] = "172.18.183.132:27017"
	ipt[19] = "172.18.183.132:27020"
	ipt[20] = "172.18.183.132:27017"
	ipt[21] = "172.18.183.132:27020"
	ipt[22] = "172.18.183.132:27017"
	ipt[23] = "172.18.183.132:27020"
	ipt[24] = "172.18.183.132:27017"
	ipt[25] = "172.18.183.132:27020"
	ipt[26] = "172.18.183.132:27017"
	ipt[27] = "172.18.183.132:27020"
	ipt[28] = "172.18.183.132:27017"
	ipt[29] = "172.18.183.132:27020"
	ipt[30] = "172.18.183.132:27017"
	ipt[31] = "172.18.183.132:27020"
	ipt[32] = "172.18.183.132:27017"
	ipt[33] = "172.18.183.132:27020"
	ipt[34] = "172.18.183.132:27017"
	ipt[35] = "172.18.183.132:27020"
	ipt[36] = "172.18.183.132:27017"
	ipt[37] = "172.18.183.132:27020"
	ipt[38] = "172.18.183.132:27017"
	ipt[39] = "172.18.183.132:27020"

	//开辟管理者协程打印收集的数据
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		n := 0
		for {
			if n == 1 && len(ch) == 0 {
				fmt.Println("管理者协程退出...")
				break
			}
			select {
			case per := <-ch:
				fmt.Println(per.Name + ":" + per.Age)
			case n = <-quit:
			}
		}
		wg.Done()
	}(wg)

	//循环开辟子协程，进行并发查询
	for i := 0; i < len(ipt); i++ {
		goch <- 1
		wg.Add(1)
		go func(ipstr string, wg *sync.WaitGroup) { //注意使用地址，进行地址传递
			fmt.Println("一个协程创建完毕")
			// str := "172.18.183.132:27017"
			session, err := mgo.Dial(ipstr)
			if err != nil {
				panic(err)
			}
			defer func() {
				<-goch
				session.Clone()
			}()

			session.SetMode(mgo.Monotonic, true)
			c := session.DB("test").C("student")

			stus := make([]Person, 500) //需大于数据库文档数
			err = c.Find(nil).All(&stus)

			fmt.Println("一共查到数据", len(stus))
			for i := 0; i < len(stus); i++ {
				ch <- stus[i]
			}
			wg.Done()
		}(ipt[i], wg)

		if i == len(ipt)-1 {
			close(goch)
			quit <- 1
		}
	}

	//等待并发子协程运行完毕后退出主协程
	wg.Wait()

	// i := 0
	// n := len(ch)
	// for ; i < n; i++ {
	// 	per := <-ch
	// 	fmt.Println(per.Name + ":" + per.Age)
	// }

	// fmt.Printf("一共%d条数据", i)
	// fmt.Println("查找完毕！")

	defer func() {
		close(ch)
		close(quit)
		fmt.Println("主协程退出")
	}()
}
