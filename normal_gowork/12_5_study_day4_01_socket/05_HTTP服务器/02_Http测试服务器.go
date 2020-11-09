package main

import (
	"fmt"
	"net/http"
)

//服务器端编写的业务逻辑处理程序
func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func main() {
	http.HandleFunc("/go", myHandler)

	//在指定的地址进行监听
	http.ListenAndServe("127.0.0.1:8000", nil)
}
