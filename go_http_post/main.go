/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/6/19 15:46
 * @File:  main
 * @Software: GoLand
 **/

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// post 请求的本质，它是request body 提交，相对于get请求(urlencoded 提交查询参数，提交内容有大小限制)
	//post的不同的形式，也就是Body的格式不同
	//post form表单是urlencoded形式，
	//post json提交的json格式
	//post 提交文件
	//post请求分为: post form,post json,post文件
	postJson()
}

func postFrom() {
	//form data形式和query string 一样
	data := make(url.Values)
	data.Add("name", "tanc")
	data.Add("age", "18")
	encode := data.Encode()

	//strings可以使用NewReader创建一个reader,因为strings包中定义了NewReader方法
	resp, _ := http.Post("http://httpbin.org/post", "application/x-www-form-urlencoded", strings.NewReader(encode))

	body, _ := io.ReadAll(resp.Body)
	defer func() { _ = resp.Body.Close() }()
	fmt.Printf("%s", body)
}

func postJson() {
	u := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		Name: "tanc",
		Age:  18,
	}

	load, _ := json.Marshal(u)
	//strings可以使用NewReader创建一个reader,因为strings包中定义了NewReader方法
	resp, _ := http.Post("http://httpbin.org/post", "application/json", bytes.NewReader(load))

	body, _ := io.ReadAll(resp.Body)
	defer func() { _ = resp.Body.Close() }()
	fmt.Printf("%s", body)
}
