package main

import "fmt"

func main() {
	var m map[int]string
	fmt.Println("m = ", m)
	fmt.Println("len(m) = ", len(m))

	mp := make(map[int]string)
	fmt.Println("mp = ", mp)
	fmt.Println("len(mp) = ", len(mp))

	mpl := make(map[int]string, 10)
	fmt.Println("mpl = ", mpl)
	fmt.Println("len(mpl) = ", len(mpl))

	mpl[0] = "赵鑫"
	mpl[1] = "boreas.zhao"
	fmt.Println("mpl = ", mpl)
	fmt.Println("len(mpl) = ", len(mpl))

}
