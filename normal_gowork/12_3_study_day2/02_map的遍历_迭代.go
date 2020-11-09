package main

import "fmt"

func main() {
	//map遍历、迭代、map的遍历时无顺序的
	m := map[int]string{1: "go", 2: "c++", 3: "python", 4: "java"}

	for key, value := range m {
		fmt.Printf("key%d ==== > value%s\n", key, value)
	}
	//判断map键值是否存在
	value, ok := m[0]
	if ok == true {
		fmt.Println("value = ", value)
	} else {
		fmt.Println("键值不存在！")
	}
	//map数值的删除，通过key 删除

	delete(m, 1)
	for key, value := range m {
		fmt.Printf("key%d ==== > value%s\n", key, value)
	}

}
