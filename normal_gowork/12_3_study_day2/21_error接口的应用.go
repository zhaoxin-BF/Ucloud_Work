package main

import "fmt"
import "errors"

func MyDiv(a, b int) (result int, err error) {
	err = nil
	if b == 0 {
		err = errors.New("分母不能为0")
	} else {
		result = a / b
	}
	return
}

func main() {
	a, b := 10, 0
	result, err := MyDiv(a, b)
	if err != nil {
		fmt.Println("err = ", err)
	} else {
		fmt.Println("结果: %d\n", result)
	}
}
