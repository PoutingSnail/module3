package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

/*
*
内容：模块二作业
完成情况：

	已完成：
		1. 完成request的header写入response的header
		2. 完成将环境变量VERSION的值写入response的header
		3. 打印客户端IP，因为不知道如何拦截所有入口，统一处理，只好在/路径和/healthz两个handleFunc里都打印
		4. 完成客户端访问/healthz的时候，返回statusCode为200
	未完成：
		对问题3， 不知道如何拦截http最终的返回response，无法get到statusCode

*
*/
func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/healthz", healthz)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// 1.将 request 中带的 header 写入 response header
func requestHeader2ResponseHeader(res http.ResponseWriter, req *http.Request) {
	for key, value := range req.Header {
		for _, v := range value {
			res.Header().Add(key, v)
		}
	}
}

// 2.读取当前系统的环境变量中的 VERSION 配置，并写入 response header
func systemVersionValue2ResponseHeader(res http.ResponseWriter) {
	version := os.Getenv("VERSION")
	fmt.Println("OS environment variable VERSION's value is:", version)
	res.Header().Set("VERSION", version)
}

// 3.Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
func foo(res http.ResponseWriter, req *http.Request) {
	if req != nil {
		fmt.Println(req.RemoteAddr)
		if req.Header != nil {
			requestHeader2ResponseHeader(res, req)
			systemVersionValue2ResponseHeader(res)
		}
	}
}

// 4.当访问 localhost/healthz 时，应返回 200
func healthz(res http.ResponseWriter, req *http.Request) {
	if req != nil {
		fmt.Println(req.RemoteAddr)
		res.WriteHeader(200)
	}
}
