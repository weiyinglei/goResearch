/*
 * array
 * 1.数组是值类型，换句话说，如果你将一个数组赋值给另外一个数组，那么，实际上就是将整个数组拷贝一份。
 * 2.如果数组作为函数的参数，那么实际传递的参数是一份数组的拷贝，而不是数组的指针。
 * 3.数组的长度也是Type的一部分，这样就说明[10]int和[20]int是两个不同的类型。
 */
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var arr1 [2]int
	arr2 := [2]int{}
	arr3 := [2]int{1, 2}
	arr3_1 := [2]int{0: 1, 1: 2}
	arr3_2 := [2]int{1: 2, 0: 1}
	arr4 := [...]int{1, 2}
	arr5 := [...]int{3, 9}
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr3_1)
	fmt.Println(arr3_2)
	fmt.Println(arr4)
	fmt.Println(arr5)
}
