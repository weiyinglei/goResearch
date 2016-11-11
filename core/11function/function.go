/*
 * function.go
 * 函数定义
 * 多值返回
 * 变参函数
 * defer
 * ？
 */
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(Add(10, 20))

	fmt.Println(divide(100, 7))
	fmt.Println(divide2(50, 7))
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(slice...))

	str, err := ReadFile("function.go")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(str)
}

func Add(a, b int) int {
	return a + b
}

func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

func divide2(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}

func sum(args ...int) int {
	s := 0
	for _, number := range args {
		s += number
	}
	return s
}

func ReadFile(strFileName string) (string, error) {
	f, err := os.Open(strFileName)
	if err != nil {
		fmt.Println()
		return "", err
	}
	defer f.Close()
	buf := make([]byte, 1024)
	var strContent string = ""
	for {
		n, _ := f.Read(buf)
		if n == 0 {
			break
		}
		strContent += string(buf[0:n])
	}
	return strContent, nil
}
