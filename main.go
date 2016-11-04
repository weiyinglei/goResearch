package main

import (
	"fmt"
	"io/ioutil"
	"github.com/weiyinglei/goResearch/test"
)

func main() {
	//文件系统http服务器
	test.Start();
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
