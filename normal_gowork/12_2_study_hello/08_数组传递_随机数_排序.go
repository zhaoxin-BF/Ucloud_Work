package main

import "fmt"
import "math/rand"
import "time"

func InitData(s *[10]int) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len(s); i++ {
		s[i] = rand.Intn(100)
	}
}

func BubbleSort(s *[10]int) {
	tmp := 0
	for i := 1; i < len(s); i++ { //控制冒泡次数
		for j := 0; j < len(s)-i; j++ { //控制每次冒泡的比较次数
			if s[j] > s[j+1] {
				tmp = s[j]
				s[j] = s[j+1]
				s[j+1] = tmp
			}
		}
	}
}

func main() {
	s := [10]int{0}
	//var a [3]int
	//fmt.Println("len(a) = ", len(a))
	InitData(&s)
	fmt.Println("s = ", s)

	BubbleSort(&s)
	fmt.Println("s_sort = ", s)
}
