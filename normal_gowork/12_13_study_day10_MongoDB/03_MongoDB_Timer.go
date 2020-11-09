package main
import (
	"fmt"
	"gopkg.in/mgo.v2"
	"time"

	//"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name string
	Phone string
}

func ManageGo(tm chan int) {
	ft := time.Now().UnixNano() / 1e6
	st := time.Now().UnixNano() / 1e6
	for {
		<-tm
		st = time.Now().UnixNano() / 1e6
		fmt.Printf("处理1000次：%v ms;\n", st-ft)
		ft = st
	}
}

func main(){
	session, err := mgo.Dial("172.18.183.132:27017")
	if err != nil{
		panic(err)
	}
	defer session.Clone()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("student")
	ft := time.Now().UnixNano() / 1e6
	st := time.Now().UnixNano() / 1e6
	for i := 0;i<1000000000000000;i++{//千万亿条
		err = c.Insert(&Person{"boreas", "666666666"})
		if err != nil{
			panic(err)
		}

		if i%100 == 0{
			st = time.Now().UnixNano() / 1e6
			fmt.Printf("插入100条数据耗时：%v ms, 共 %v 条数据 ;\n", st-ft, i)
			ft = st
		}
	}

	fmt.Println("插入完毕！")
}




