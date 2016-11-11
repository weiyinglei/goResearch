/*
 * defineVar.go
 */
package main

import "fmt"

var x, y, z int
var f float32 = 1.6
var s, n = "abc", 123
var (
	a int
	b float32
)

func main() {
	fmt.Println(x)
	x := 123
	fmt.Println(x)
	fmt.Println(x, s, n)
}
