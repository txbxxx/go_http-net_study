/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/6/19 15:08
 * @File:  main
 * @Software: GoLand
 **/

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Reader 定义一个自己的Reader
type Reader struct {
	io.Reader
	Total   int64
	Current int64
}

// 实现Reader接口
func (r *Reader) Read(p []byte) (n int, err error) {
	n, err = r.Reader.Read(p)
	r.Current += int64(n)
	// 格式化进度条
	fmt.Printf("\r 进度: %.2f%% \n", float64(r.Current*10000/r.Total)/100)

	return
}

func main() {
	//自动文件下载、图片下载、压缩包下载
	//测试下载go安装包
	const url = "https://golang.google.cn/dl/go1.22.4.windows-amd64.msi"
	//下载到的文件命名
	const filename = "go1.22.4.msi"
	downloadProgress(url, filename)
}

// 普通下载下载
func download(url, filename string) {
	// 使用Get请求获取数据
	r, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer func() { _ = r.Body.Close() }()
	// 创建文件，以便后面拷贝数据
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer func() { _ = f.Close() }()
	//将获取倒的数据拷贝到文件中
	n, err := io.Copy(f, r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
}

// 带进度条的下载
func downloadProgress(url, filename string) {
	// 使用Get请求获取数据
	r, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer func() { _ = r.Body.Close() }()
	// 创建文件，以便后面拷贝数据
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer func() { _ = f.Close() }()

	reader := &Reader{
		// 通过r.Body获取数据
		Reader: r.Body,
		// 通过r.ContentLength获取文件大小
		Total: r.ContentLength,
	}

	//将获取倒的数据拷贝到文件中
	n, err := io.Copy(f, reader)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
}
