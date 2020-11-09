package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s := a[2:5]
	s1 := s[3:8]

	//s[1] = 666
	fmt.Println("s = ", s)
	fmt.Println("s1 = ", s1)
	fmt.Println("a = ", a)
	fmt.Printf("len(s) = %d, cap(s) = %d\n", len(s), cap(s))

}
