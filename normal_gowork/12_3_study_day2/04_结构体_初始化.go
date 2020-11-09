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
	//顺序初始化，每一个都得初始化
	var s1 Student = Student{1, "mike", 'm', 22, "Shanghai"}

	fmt.Println("s1 = ", s1)

	fmt.Printf(" id = %d\n name = %s\n sex = %c\n age = %d\n addr = %s\n",
		s1.id, s1.name, s1.sex, s1.age, s1.addr)

	//指定成员初始化的，会自动初始化为0，string类型不初始化
	var s2 Student = Student{id: 2, name: "boreas"}
	fmt.Println("s2 = ", s2)

	//指针类型顺序初始化，必须取地址
	var sp1 *Student = &Student{1, "mike", 'm', 22, "Shanghai"}
	fmt.Println("*sp1 = ", *sp1)
	//fmt.Println("*sp1 = ", sp1)
	//指针类型也可以直接打印，
	//因为go语言在内部进行了转换，但是打印出的为取地址的&{*}

	//指针类型指定成员初始化，同样必须取地址
	var sp2 *Student = &Student{id: 2, name: "boreas"}
	fmt.Println("*sp2 = ", *sp2)
}
