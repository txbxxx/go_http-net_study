/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/6/19 10:42
 * @File:  main
 * @Software: GoLand
 **/

package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"
	"io"
	"net/http"
)

func main() {
	r, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println(err)
	}
	defer func() { _ = r.Body.Close() }()
	//body(r)
	//status(r)
	//header(r)
	encoding(r)

}

// 响应体
func body(r *http.Response) {
	//获取响应体
	Body, _ := io.ReadAll(r.Body)
	fmt.Println(string(Body))
}

// 状态码
func status(r *http.Response) {
	//状态码
	fmt.Println(r.StatusCode)
	//状态描述信息
	fmt.Println(r.Status)

}

// 响应头
func header(r *http.Response) {
	//获取响应头信息
	fmt.Println(r.Header.Get("Content-Type"))
	//r.Header底层类型是一个map
	//不推荐这种方法因为Get会忽略大小写，用map则不用
	fmt.Println(r.Header["Content-Type"])

}

// 编码
func encoding(r *http.Response) {
	// content-type 会提供编码信息
	// 可以通过html head mate 获取编码信息
	// 可以通过网页头部猜测编码信息
	//使用DetermineEncoding
	bufReader := bufio.NewReader(r.Body)
	//使用这个缓冲读取器从响应体中预读（不移动读取位置）最多 1024 个字节的数据，并将这些数据存储到 body 字节切片中。这样可以在不实际消耗这些数据的情况下提前查看一部分响应体的内容。
	body, _ := bufReader.Peek(1024)
	//它返回三个参数分别是：文件编码结构体、编码名和是否确定是这个编码，它接收两个个参数：字节切片（就是这个响应体的1024个字节）和响应头信息。
	e, _, _ := charset.DetermineEncoding(body, r.Header.Get("Content-Type"))
	//如果编码不同需要执行编码转换
	reader := transform.NewReader(bufReader, e.NewDecoder())
	all, _ := io.ReadAll(reader)

	fmt.Println(string(all))
}
