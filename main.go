package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
)

func main() {
	//文件系统http服务器
	StartFileServer();
}

func main1() {
	bytes,err := ioutil.ReadFile("d:/1.txt")
	if err == nil {
		fmt.Println(bytes)
		fmt.Println(string(bytes))
		println(bytes)
		println(string(bytes))
	}
}

func StartFileServer() {
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("d:/")))) // 正确
	//http.Handle("/", http.FileServer(http.Dir("public"))) // 正确（访问根目录时转到public目录）

	//http.Handle("/public", http.StripPrefix("/public", http.FileServer(http.Dir("public")))) // 错误
	//http.Handle("/public", http.FileServer(http.Dir("/public"))) // 错误
	//http.Handle("/public", http.FileServer(http.Dir("/public/"))) // 错误
	//http.Handle("/public", http.FileServer(http.Dir("./public"))) // 错误！
	log.Fatal(http.ListenAndServe(":8080", nil))
}