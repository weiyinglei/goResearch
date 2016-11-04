package test

import (
	"net/http"
	"log"
)

func Start() {
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("d:/")))) // 正确
	//http.Handle("/", http.FileServer(http.Dir("public"))) // 正确（访问根目录时转到public目录）

	//http.Handle("/public", http.StripPrefix("/public", http.FileServer(http.Dir("public")))) // 错误
	//http.Handle("/public", http.FileServer(http.Dir("/public"))) // 错误
	//http.Handle("/public", http.FileServer(http.Dir("/public/"))) // 错误
	//http.Handle("/public", http.FileServer(http.Dir("./public"))) // 错误！
	log.Fatal(http.ListenAndServe(":8080", nil))
}