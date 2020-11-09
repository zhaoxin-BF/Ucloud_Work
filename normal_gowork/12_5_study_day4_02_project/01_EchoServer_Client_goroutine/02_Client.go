package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		conn.Close()
		fmt.Println("conn closed...")
	}()
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)

	for {
		time.Sleep(7 * time.Second)
		fmt.Println("i am still alive..")
	} //这里通过增加一个for循环，防止main的goroutine退出，一旦main退出了，其它子goroutine会全部退出;
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
