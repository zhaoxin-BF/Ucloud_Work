package main

import (
	"fmt"
	"net"
	"strings"
)

//1、处理连接请求的函数
func HandleConn(conn net.Conn) {
	//函数调用完毕后，自动关闭conn
	defer conn.Close()

	//获取客户端的网络地址信息
	addr := conn.RemoteAddr().String()
	fmt.Printf("用户[%s]: connect successed\n", addr)

	//定义接收切片，用以接受数据
	buf := make([]byte, 2048)
	for {
		//读取用户数据
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("err = ", err)
			return
		}
		fmt.Printf("收到用户[%s]数据：%+v\n", addr, string(buf[:n-2]))

		//fmt.Println("len = ", len(buf[:n-2])) //除去末尾换行/n/t两个转义字符
		//检测用户输入为exit后推出断开连接
		if "exit" == string(buf[:n-2]) {
			fmt.Println(addr, "exit")
			return
		}
		//把数据转换为大写， 发送给用户
		conn.Write([]byte(strings.ToUpper(string(buf[:n-1]))))
	}

}

func main() {
	//1、监听端口
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	//2、设置程序退出后关闭监听端口
	defer listener.Close()

	//接受到用户请求后，创建多个协程处理请求
	for {
		conn, err := listener.Accept() //3、阻塞等待接受请求
		if err != nil {                //判断出错
			fmt.Println("err = ", err)
			continue
		}
		//4、处理用户请求, 新建一个协程, 并传入连接句柄给函数
		go HandleConn(conn)
	}

}
