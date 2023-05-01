package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

//接收客户端 request，并将 request 中带的 header 写入 response header
func readHeader(w http.ResponseWriter, r *http.Request) {
	log.Println("header info: ", r.Header)
	for k, v := range r.Header {
		w.Header().Set(k, v[0])
		log.Println("header info: ", fmt.Sprintf("%s=%s\n", k, v))
	}
}

//读取当前系统的环境变量中的 VERSION 配置，并写入 response header
func readVersion(w http.ResponseWriter, r *http.Request) {
	//读取env中的VERSION
	version := os.Getenv("VERSION")
	// io.WriteString(w, fmt.Sprintf("%s=%s", Version, version))
	w.Header().Set("Version", version)
	log.Println("header info: ", version)
}

//Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
func getClientinfo(w http.ResponseWriter, r *http.Request) {
	form := r.RemoteAddr
	log.Println("Client->ip:port=" + form)
	ipStr := strings.Split(form, ":")
	log.Println("Client->ip=" + ipStr[0]) //打印ip
	// 获取http响应码
	println("Client->response code=" + strconv.Itoa(http.StatusOK))
	//println("response code->：" + code)
	io.WriteString(w, "succeed")

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/readHeader", readHeader)
	mux.HandleFunc("/readVersion", readVersion)
	mux.HandleFunc("/clientinfo", getClientinfo)
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "OK")
	})
	err := http.ListenAndServe(":80", mux)
	if err != nil {
		log.Fatal(err)
	}
}
