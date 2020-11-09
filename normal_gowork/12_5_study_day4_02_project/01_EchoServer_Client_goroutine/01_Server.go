package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handConn(conn) //启动一个goroutine处理当前连接请求
	}
}

func handConn(conn net.Conn) {
	defer func() {
		conn.Close()
	}()

	input := bufio.NewScanner(conn)
	for input.Scan() {
		go echo(conn, input.Text(), 1*time.Second)
	}
}

func echo(conn net.Conn, text string, delay time.Duration) {
	fmt.Fprintln(conn, "\t", strings.ToUpper(text))
	time.Sleep(delay)
	fmt.Fprintln(conn, "\t", text)
	time.Sleep(delay)
	fmt.Fprintln(conn, "\t", strings.ToLower(text))
}
