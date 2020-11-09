package main

import "fmt"

func testa() {
	fmt.Println("aaaaaaaaaa")
}
func testb(x int) {
	var a [10]int
	a[x] = 111 //当x为=20的时候，导致数组越界，产生一个panic//内部painc
}
func testc() {
	fmt.Println("ccccccccccc")
}

func main() {
	testa()
	testb(20)
	testc()
}
