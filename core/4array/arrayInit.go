/*
 * array
 * 1.数组是值类型，换句话说，如果你将一个数组赋值给另外一个数组，那么，实际上就是将整个数组拷贝一份。
 * 2.如果数组作为函数的参数，那么实际传递的参数是一份数组的拷贝，而不是数组的指针。
 * 3.数组的长度也是Type的一部分，这样就说明[10]int和[20]int是两个不同的类型。
 */
package main

import (
	"fmt"
	//"reflect"
)

//演示array初始化
func main() {
	var arr11 [2]int //所有值都初始化为0
	arr12 := [2]int{} //所有值都初始化为0
	arr21 := [3]int{1, 2, 3}
	arr22 := [3]int{0: 1, 1: 2, 2: 3}
	arr23 := [3]int{2: 3, 1: 2, 0: 1}
	arr24 := [...]int{1, 2, 3}

	fmt.Println("arr11 is ", arr11)
	fmt.Println("arr12 is ", arr12)
	fmt.Println("arr21 is ", arr21)
	fmt.Println("arr22 is ", arr22)
	fmt.Println("arr23 is ", arr23)
	fmt.Println("arr24 is ", arr24)
	fmt.Println("arr24[1] is ", arr24[1])
	fmt.Println("arr24 length is ", len(arr24))
}
