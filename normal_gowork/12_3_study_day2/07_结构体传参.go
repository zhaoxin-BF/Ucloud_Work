package main

import "fmt"

type Student struct {
	id   int
	name string
	sex  byte //字符型
	age  int
	addr string
}

//结构体做参数，是传值的，没办法在函数体内修改值
func test01(s Student) {
	s.id = 10
	fmt.Println("test s = ", s)
}

func test02(s *Student) { //值传递，形参无法改实参，所以做参数得传指针
	s.id = 100
}

func main() {
	s := Student{1, "mike", 'm', 22, "sh"}

	test01(s)

	fmt.Println("main s = ", s)

	test02(&s)
	fmt.Println("test02 s = ", s)

}
