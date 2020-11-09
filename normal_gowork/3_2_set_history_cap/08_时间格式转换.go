package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now().Unix()
	st := time.Unix(t,0).Format("2006-01-02")
	fmt.Println(st)
}

