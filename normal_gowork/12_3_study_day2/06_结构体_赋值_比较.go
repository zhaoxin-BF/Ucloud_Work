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
	s1 := Student{1, "mike", 'm', 18, "sh"}
	s2 := Student{1, "mike", 'm', 18, "sh"}
	s3 := Student{3, "mike", 'm', 18, "sh"}

	fmt.Println("s1 == s2", s1 == s2)
	fmt.Println("s1 == s3", s2 == s3)

	s3 = s2
	fmt.Println("s3 = ", s3)

	var tmp Student
	tmp = s3
	fmt.Println("tmp = ", tmp)
}
