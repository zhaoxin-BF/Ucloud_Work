package main

import "fmt"

type Humaner interface {
	//方法，只有声明，没有实现，由别的类型（自定义）实现
	sayhi()
}

type Student struct {
	name string
	id   int
}

//Teacher实现了此方法
func (tmp *Student) sayhi() {
	fmt.Printf("Student[%s, %d] sayhi\n", tmp.name, tmp.id)
}

type Teacher struct {
	addr  string
	group string
}

//Teacher实现了此方法
func (tmp *Teacher) sayhi() {
	fmt.Printf("Teacher[%s, %d] sayhi\n", tmp.addr, tmp.group)
}

type MyStr string

//MyStr实现了此方法
func (tmp *MyStr) sayhi() {
	fmt.Printf("MyStr[%s] sayhi\n", *tmp)

}

//类比C++ 的多态

func WhoSayHi(i Humaner) {
	i.sayhi()
}

func main() {
	s := &Student{"mike", 666}      //指针
	t := &Teacher{"shanghai", "go"} //指针
	var str MyStr = "Hello mike"

	WhoSayHi(s)
	WhoSayHi(t)
	WhoSayHi(&str) //传的指针

	//创建一个切片
	x := make([]Humaner, 3)
	x[0] = s
	x[1] = t
	x[2] = &str

	for _, i := range x {
		i.sayhi()
	}
}

// func main() {
// 	//定义接口类型的变量
// 	var i Humaner

// 	//只要实现了这个方法的类型，
// 	s := &Student{"mike", 666}
// 	i = s
// 	i.sayhi()

// 	t := &Teacher{"shanghai", "go"}
// 	i = t
// 	i.sayhi()

// 	var str MyStr = "Hello mike"
// 	i = &str
// 	i.sayhi()
// }
