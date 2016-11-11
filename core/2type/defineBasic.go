//define.go
package main

import "fmt"

func main() {
	var b bool
	var n int     /*定义单个变量*/
	var i int = 3 /*定义单个变量并赋值*/
	var (         /*多变量的定义*/
		aa  int = 3
		str string
	)
	var i1, i2, i3 int = 1, 2, 3 /*定义多个变量并赋值*/
	var strName = "张三"           /*Go会自动检测变量的类型*/
	strSex := "男"                /*：=定义变量，并给变量赋值，可以省略var关键字*/
	fmt.Println(b, n, i, aa, str, i1, i2, i3, strName, strSex)
}
