/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/6/18 16:08
 * @File:  main
 * @Software: GoLand
 **/

package main

import (
	"bytes"
	json2 "encoding/json"
	"fmt"
	"io"
	"net/http"
)

//func Index(writer http.ResponseWriter, request *http.Request) {
//	//请求路径
//	fmt.Println(request.Method, request.URL.String())
//	//获取请求体
//	if request.Method != "GET" {
//		byteData, _ := io.ReadAll(request.Body)
//		fmt.Println(string(byteData))
//	}
//	//获取请求头
//	fmt.Println(request.Header)
//
//	//响应请求
//	writer.Write([]byte("hello world!!"))
//}

type Data struct {
	username string
	password string
}

func main() {
	Delete()

}

func Delete() {
	//新建一个请求
	r, _ := http.NewRequest(http.MethodDelete, "http://httpbin.org/delete", nil)

	//将请求通过Client的Do方法发送出去
	put, err := http.DefaultClient.Do(r)
	if err != nil {
		fmt.Println(err)
	}
	defer put.Body.Close()
	Body, err := io.ReadAll(put.Body)
	fmt.Println(string(Body))
}

func Put() {
	//Put示例

	//新建一个请求
	r, _ := http.NewRequest(http.MethodPut, "http://httpbin.org/put", nil)

	//将请求通过Client的Do方法发送出去
	put, err := http.DefaultClient.Do(r)
	if err != nil {
		fmt.Println(err)
	}
	defer put.Body.Close()
	Body, err := io.ReadAll(put.Body)
	fmt.Println(string(Body))

}

func Post() {
	//Post示例
	//请求数据
	data := Data{
		username: "tanc",
		password: "123456",
	}
	//转换为json格式
	json, _ := json2.Marshal(data)
	//读取json格式,bytes包中实现了io.Reader接口
	jsonReader := bytes.NewReader(json)

	r, err := http.Post("https://spafter.tanc.fun:9999/loginUser", "application/json", jsonReader)
	if err != nil {
		fmt.Println(err)
	}
	Body, err := io.ReadAll(r.Body)
	fmt.Println(string(Body))
	defer r.Body.Close()
}

func Get() {
	//创建Get请求
	r, err := http.Get("https://baidu.com")
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	fmt.Println(string(body))
}
