//1、监听端口
//2、设置推出后关闭监听
//3、循环等待客户端连接
//4、为每个客户端创建一个协程处理请求

//5、请求处理函数的编写

package main

import (
	"fmt"
	"net"
	"runtime"
	"strings"
)

//五、请求处理函数
func Handler(conn net.Conn) {
	//1、退出后关闭连接
	defer conn.Close()

	//获取客户端的网络地址信息，IP地址
	addr := conn.RemoteAddr().String()
	fmt.Printf("用户[%s] connect successed\n", addr)

	//2、创建接受的切边（简单来说就是先定义接受的空间）,重点是byte类型，且需要定义大小
	buf := make([]byte, 2048)

	//3、循环阻塞接受数据并回传给客户端
	for {
		//4、收到数据，n为收到数据的大小 单位为byte
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Read err: ", err)
			return //出错退出连接
		}

		//将收到的数据打印在屏幕上
		fmt.Printf("用户[%s]:%s\n", addr, string(buf[:n]))

		//5、检测用户输入为exit后，退出并断开连接
		if "exit" == string(buf[:n-1]) { //出去数据末尾的换行及火车符
			fmt.Printf("用户%s Exit\n", addr)
			return
		}

		//6、处理接受到的数据，小写转大写，并发送给用户
		// conn.Write([]byte(strings.ToUpper(string(buf[:n-1]))))
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}
}

func main() {
	runtime.GOMAXPROCS(8)
	//1、监听端口
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Listen err: ", err)
		return
	}
	fmt.Println("服务器开启成功！")

	//2、设置推出后关闭监听
	defer listener.Close()
	//3、循环等待客户端连接
	n := 0
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept err: ", err)
			continue
		}

		//4、为每个客户端创建一个协程处理请求
		go Handler(conn)
		n++
		if n%200 == 0 {
			fmt.Println("一共有 ", n, " 个客户端连接")
		}
	}

}
