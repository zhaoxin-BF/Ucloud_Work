package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func WriteFile(path string) {
	//打开文件，新建文件
	f, err := os.Create(path) //f为文件指针，文件描述符
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	//使用完毕，需要关闭文件
	defer f.Close() //defer的作用是，在函数执行完毕后再调用这一关闭操作

	var buf string
	for i := 0; i < 10; i++ {
		//格式化完后，存贮到返回值
		buf = fmt.Sprintf("i  = %d \n", i)

		n, err := f.WriteString(buf)
		if err != nil {
			fmt.Println("err = ", err)
			return
		}
		fmt.Println("n = ", n)
	}
}

//每次读取一行
func ReadFileLine(path string) {
	//打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	//关闭文件
	defer f.Close()

	//新建一个缓冲区，把内容先放在缓冲区
	r := bufio.NewReader(f)

	for {
		//遇到'\n'结束读取,但是\n也会被读取进去
		buf, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF { //文件已经结束
				break
			}
			fmt.Println("err = ", err)
		}
		fmt.Printf("buf = #%s#\n", string(buf))
	}

}

//每次读取全部，buf 大小的文件内容
func ReadFile(path string) {
	//打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	//关闭文件
	defer f.Close()

	buf := make([]byte, 1024*2) //2K大小

	//n代表从为文件读取内容的长度
	n, err1 := f.Read(buf)
	if err1 != nil && err1 != io.EOF { //文件出错，同时没到结尾
		fmt.Println("err1 = ", err1)
		return
	}

	fmt.Printf("内容：\n%v", string(buf[:n])) //切片
}
func main() {
	path := "./demo.txt"
	WriteFile(path)
	//
	//ReadFile(path)
	//ReadFileLine(path)
}
