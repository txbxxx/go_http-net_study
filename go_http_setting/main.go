/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/6/18 21:36
 * @File:  main.go
 * @Software: GoLand
 **/

package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func printBody(r *http.Response) {
	Body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()
	fmt.Printf("%s", Body)
}

func main() {
	//设置请求查询参数比如http://httpbin.roge/get?name=tanchang&age=18
	//Params()
	//定制请求头
	Header()
}

func Header() {
	request, err := http.NewRequest(http.MethodGet, "http://httpbin.org/get", nil)
	if err != nil {
		fmt.Println(err)
	}
	//添加请求头,可以通过User-agent进行反爬虫
	request.Header.Add("user-agent", "chrome")
	//发起请求
	r, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	//打印Body
	printBody(r)
}

func Params() {
	request, err := http.NewRequest(http.MethodGet, "http://httpbin.org/get", nil)
	if err != nil {
		fmt.Println(err)
	}
	//url.Values 类型可以用来组织Get请求的参数，它是一个map String类型的别名
	params := make(url.Values)
	//这里就相当于get?name=tanc&password=123455
	params.Add("name", "tanc")
	params.Add("password", "123455")

	request.URL.RawQuery = params.Encode()

	//发起请求
	r, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	//打印Body
	printBody(r)
}
