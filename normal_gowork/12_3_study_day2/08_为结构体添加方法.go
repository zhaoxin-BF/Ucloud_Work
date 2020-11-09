package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

//带有接受者的函数叫做方法、普通类型（tmp Persion）为接收者
func (tmp Person) PrintInfo() {
	fmt.Println("tmp = ", tmp)
}

//成员函数、指针类型，接受者类型本身不能是指针类型///修改的时候需要传指针
func (p *Person) SetInfo(n string, s byte, a int) {
	p.name = n
	p.sex = s
	p.age = a
}

func main() {
	//定义同时初始化
	p := Person{"mike", 'm', 18}
	p.PrintInfo()

	//定义一个结构体变量
	var p2 Person
	(&p2).SetInfo("yoyo", 'f', 22)
	p2.PrintInfo()
}

//函数重载、只与接受这类型有关
func (tmp int) test() {

}

func (tmp byte) test() {

}
