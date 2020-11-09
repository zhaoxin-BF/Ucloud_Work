package main

import "fmt"
import "math/rand"
import "time"

func InitData(s []int) {
	//设置种子
	rand.Seed(time.Now().UnixNano()) //以动态时间设置随机数
	for i := 0; i < len(s); i++ {
		s[i] = rand.Intn(100) //100以内的随机数
	}
}

func BubbleSort(s []int) {
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
	n := 10

	s := make([]int, n)

	InitData(s) //初始化（切片）数组
	fmt.Println("排序前", s)

	BubbleSort(s)
	fmt.Println("排序后", s)

}
