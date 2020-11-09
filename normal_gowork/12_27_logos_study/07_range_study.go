package main

import (
	"fmt"
)

func modifySlice() {
	data := map[string]map[string]string{"a":{"a": "A"}, "b":{"b": "B"}, "c":{"c": "C"}}
	for _, v := range data {
		fmt.Println("modify Mapping", v)
	}
}

func main() {
	modifySlice()
}