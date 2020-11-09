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
	p := Person{"mike", 'm', 18} //普通变量
	p.SetInfoPointer()           //内部转换
}
