/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/6/19 21:12
 * @File:  main
 * @Software: GoLand
 **/

package main

import (
	"fmt"
	"net/http"
)

func main() {
	////注册路由
	//http.Handle("/index", &MyHandler{})
	////监听端口
	//http.ListenAndServe(":8080", nil)

	http.HandleFunc("/index", func(responseWriter http.ResponseWriter, request *http.Request) {
		fmt.Println(responseWriter, "index")
	})
	http.ListenAndServe(":8080", nil)
}

// MyHandler 创建自定义Handler
type MyHandler struct {
}

// 实现Handler接口
func (h *MyHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("my implement")
}
