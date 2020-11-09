package main

import "fmt"

func main() {
	s := make([]int, 0, 1)
	fmt.Printf("len(s) = %d, cap(s) = %d\n", len(s), cap(s))
	oldCap := cap(s)

	for i := 0; i < 20; i++ {
		s = append(s, i)
		if newCap := cap(s); oldCap < newCap {
			fmt.Printf("cap %d ==> %d\n", oldCap, newCap)
			oldCap = newCap
		}
	}
}
