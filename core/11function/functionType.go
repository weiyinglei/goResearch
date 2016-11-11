/*
 * functionType.go
 * ？
 */
package main

import "fmt"

type MyFuncType func(int) bool

func IsBigThan5(n int) bool {
	return n > 5
}

func Display(arr []int, f MyFuncType) {
	for _, v := range arr {
		if f(v) {
			fmt.Println(v)
		}
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	Display(arr, IsBigThan5)
}
