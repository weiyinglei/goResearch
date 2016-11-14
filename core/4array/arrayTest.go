package main

import (
	"fmt"
)

/*
 * 声明及赋值
 * 比较
 * 使用new创建数组 //TODO
 * n维数组
 * 冒泡排序 : 除了冒泡排序还有快速排序之类的 //TODO
 */
func main() {
	defineArray()
	twoArray()
	nArray()
	bubble()
}

func defineArray() {
	a := [10]int{9: 1}
	fmt.Println(a)

	b := [10]int{
		1: 10,
		2: 5,
		3: 2,
	}
	fmt.Println(b)
}

func twoArray() {
	a := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(a)
}

func nArray() {
	a := [2][3][4]int{
		{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
		},
		{
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		},
	}
	fmt.Println(a)
}

func bubble() {
	a := [...]int{6, 3, 1, 5, 9}
	fmt.Println(a)

	num := len(a)
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if a[j] > a[i] {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
	fmt.Println(a)
}
