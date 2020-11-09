package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

func (p *Person) SetInfoValue() {
	fmt.Println("SetInfoValue")
}

func (p Person) SetInfoPointer() {
	fmt.Println("SetInfoPointer")
}

func main() {
	//结构体变量是一个指针变量， 他能够调用的方法，叫做方法的集合 ，简称方法集
	p := &Person{"mike", 'm', 18} //指针变量

	p.SetInfoPointer()
	p.SetInfoValue() //（*p）内部做的转化，指针P转成*p

	(*p).SetInfoValue()
	(*p).SetInfoPointer() //同内部转换
}
