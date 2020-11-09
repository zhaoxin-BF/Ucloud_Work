package main

import (
	"fmt"
	"net"
)

func main() {
	//监听
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("Listen err = ", err)
		return
	}

	fmt.Println("服务器启动完成!")
	defer listener.Close()

	//阻塞等待用户连接
	conn, err1 := listener.Accept()
	if err1 != nil {
		fmt.Println("listener.Accept err = ", err1)
		return
	}

	//接受用户数据
	buf := make([]byte, 1024)
	n, err2 := conn.Read(buf)
	if n == 0 {
		fmt.Println("read err = ", err2)
		return
	}

	fmt.Printf("#%s#", string(buf[:n]))
}
