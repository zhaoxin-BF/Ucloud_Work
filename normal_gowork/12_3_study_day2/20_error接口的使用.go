package main

import "fmt"
import "errors"

func main() {
	//fmt.Errorf("%s", "this is normol err")
	err1 := fmt.Errorf("%s", "this is normol err1")
	fmt.Println("err1 = ", err1)

	err2 := errors.New("This is normal err2")
	fmt.Println("err2 = ", err2)
}
