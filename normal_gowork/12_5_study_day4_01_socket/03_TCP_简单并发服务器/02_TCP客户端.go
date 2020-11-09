package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//1.主动连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Dial err = ", err)
		return
	}

	//2.main调用完毕， 关闭连接
	defer conn.Close()

	//接收服务器回复的数据，输出到屏幕
	go func() {
		//切片缓冲
		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf) //接收服务器的内容
			if err != nil {
				fmt.Println("conn.Read err = ", err)
				return
			}
			fmt.Println("收到服务器返回消息：", string(buf[:n]))
		}
	}()

	//从键盘读取内容，给服务器发送内容
	str := make([]byte, 1024)
	for {
		n, err := os.Stdin.Read(str) //从键盘读取内容， 放在str中
		if err != nil {
			fmt.Println("os.Stdin.Read err =", err)
			return
		}

		//把输入的内容给服务器发送
		conn.Write(str[:n])
	}
}
