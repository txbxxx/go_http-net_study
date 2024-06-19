/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/6/19 16:29
 * @File:  main
 * @Software: GoLand
 **/

package main

import (
	"errors"
	"fmt"
	"net/http"
)

func main() {
	redirectForbidden()
}

func redirectLimet() {
	//限制重定向次数
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if len(via) > 10 {
				return errors.New("重定次数太多")
			}
			return nil
		},
	}
	r, _ := http.NewRequest(http.MethodGet, "http://httpbin.org/redirect/20", nil)

	_, err := client.Do(r)
	if err != nil {
		panic(err)
	}
}

func redirectForbidden() {
	//禁止重定向
	// 登录请求，防止重定向到首页
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	r, _ := http.NewRequest(http.MethodGet, "http://httpbin.org/cookies/set?name=tanc", nil)

	//rtwo, err := http.DefaultClient.Do(r)
	rtwo, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	defer func() { _ = rtwo.Body.Close() }()
	fmt.Println(rtwo.Request.URL)

}
