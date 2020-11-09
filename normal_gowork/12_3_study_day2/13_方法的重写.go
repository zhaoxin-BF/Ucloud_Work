package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

func (p *Person) SetInfoValue() {
	fmt.Printf("name = %s, sex=%c, age=%d\n", p.name, p.sex, p.age)
}

//学生结构体， 继承了Person字段， 成员和方法都继承了
type Student struct {
	Person //匿名字段
	id     int
	addr   string
}

func (p *Student) SetInfoValue() {
	fmt.Println("Student = ", p)
}

func main() {
	s := Student{Person{"mike", 'm', 18}, 1, "shanghai"}
	s.SetInfoValue()

	s.Person.SetInfoValue()

}
