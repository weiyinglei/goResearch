/*
 * defineSlice.go
 */
package main

import (
	"fmt"
	"reflect"
)

func main() {
	p := [...]int{2, 3, 5, 7, 11, 13} //定义一个数组
	s1 := p[1:3]                      //定义一个切片，包含3,5两个元素
	s2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(p)
	fmt.Println(s1)
	fmt.Println(reflect.TypeOf(p))
	fmt.Println(reflect.TypeOf(s1))
	changeArrayValue(p)
	fmt.Println(p)
	changeSliceValue(s1)
	fmt.Println(s1)
	fmt.Println(p)
	fmt.Println(reflect.TypeOf(s2))
}

func changeArrayValue(arr [6]int) {
	arr[0] = 100
}

func changeSliceValue(slice []int) {
	slice[0] = 100
}
