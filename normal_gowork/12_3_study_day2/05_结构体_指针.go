package main

import "fmt"

type Student struct {
	id   int
	name string
	sex  byte //字符型
	age  int
	addr string
}

func main() {
	//1、指针有合法指向后，才操成员
	//2、先定义一个普通结构体变量
	var s Student

	//定义一个指针变量，保存s的地址
	var p1 *Student
	p1 = &s

	//通过指针操作成员， p1.id 和(*p1).id 完全等价

	p1.id = 18
	(*p1).name = "mike"
	p1.sex = 'm'

	fmt.Println("p1 = ", p1)
	fmt.Println("*p1 = ", *p1)

	//方式二 通过new来进行操作，返回一个结构体的指针
	p2 := new(Student)

	p2.name = "boreas"
	fmt.Println("p2 = ", p2)
}
