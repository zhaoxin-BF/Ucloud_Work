package main

import "fmt"

func Sum(n int) {

	if n == 1 {
		fmt.Printf("number = %d\n", 1)
	} else {
		Sum(n - 1)
		fmt.Printf("number = %d\n", n)
	}
}

func main() {
	Sum(100)
}
