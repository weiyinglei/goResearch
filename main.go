package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
	"os/exec"
	"bytes"
)

func main() {
	//文件系统http服务器
	//startFileServer("/public/","d:/");
	//读取文件
	//readFile("d:/1.txt")
	//system("ss-bash/ssadmin.sh show 10001 ") //调用函数，参数是指令，可以有多条
}

//调用系统指令的方法，参数s 就是调用的shell命令
func system(s string) {
	cmd := exec.Command("/bin/sh", "-c", s) //调用Command函数
	var out bytes.Buffer //缓冲字节

	cmd.Stdout = &out //标准输出
	err := cmd.Run() //运行指令 ，做判断
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", out.String()) //输出执行结果
}

func readFile(fileName string) {
	bytes,err := ioutil.ReadFile(fileName)
	if err == nil {
		fmt.Println(bytes)
		fmt.Println(string(bytes))
		println(bytes)
		println(string(bytes))
	}
}

func startFileServer(url string, filePath string) {
	http.Handle(url, http.StripPrefix(url, http.FileServer(http.Dir(filePath)))) // 正确
	//http.Handle("/", http.FileServer(http.Dir("public"))) // 正确（访问根目录时转到public目录）

	//http.Handle("/public", http.StripPrefix("/public", http.FileServer(http.Dir("public")))) // 错误
	//http.Handle("/public", http.FileServer(http.Dir("/public"))) // 错误
	//http.Handle("/public", http.FileServer(http.Dir("/public/"))) // 错误
	//http.Handle("/public", http.FileServer(http.Dir("./public"))) // 错误！
	log.Fatal(http.ListenAndServe(":8080", nil))
}