package main

import (
	"fmt"
	"net"
	"runtime"
	// "time"
)

func Client(num int) {
	//1、主动连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Dial err: ", err)
		return
	}

	//2、main调用完毕关闭连接
	defer conn.Close()

	//3、发送和接受服务器数据，并在桌面输出
	str := make([]byte, 1024) //创建发送给服务器的数据
	buf := make([]byte, 1024) //接受服务器返回的数据
	for {
		//1、发送消息
		str = []byte("hello world")

		n := len(str)
		conn.Write(str[:n])

		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Read err: ", err)
		}
		// fmt.Println("服务器回复：", string(buf[:bn]))
	}
}

func main() {
	runtime.GOMAXPROCS(8)
	for i := 0; i <= 3; i++ {
		go Client(i)
	}
	fmt.Println("客户端连接请求发送完毕！")

	for {
		//控制主协程不要退出
	}
}
