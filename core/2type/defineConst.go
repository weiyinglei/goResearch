/*
 * defineConst.go
 */
package main

import "fmt"

func main() {
	const PI = 3.1415926
	const y = "Hello"
	const (
		z = false
		a = 123
	)
	const username, sex = "张三", "男"
	const (
		Sunday, aa  = iota, iota * 2
		Monday, bb  = 1, 1
		Tuesday, cc = iota, iota * 2
	)
	fmt.Println(PI, y, z, a, username, sex)
	fmt.Println(Sunday, Monday, Tuesday)
	fmt.Println(aa, bb, cc)
}
