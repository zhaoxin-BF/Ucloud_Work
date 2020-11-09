package main

import "fmt"

func max(a, b int) (max, min int) {
	if a > b {
		max = a
		min = b
	} else {
		max = b
		min = a
	}
	return
}

func main() {
	max, min := max(300, 200)
	fmt.Printf("Max = %d, Min = %d", max, min)
}
