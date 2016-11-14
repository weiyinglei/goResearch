package main

import (
	"fmt"
)

/*
 * 参数列表
 * 返回值 命名返回值，不命名返回值
 * 不定长变参
 * 值类型，引用类型参数
 * 匿名函数
 * 闭包
 * defer
 */
func main() {
	E(1, "test interface", func() {
		fmt.Println("Hello World")
	})
	// F()
	// f := closure(1)
	// fmt.Println(f(1))
	// defer1()
	// defer2()
	// defer3()
	// p1()
	// panic1()
	// p2()
	// test()
}

//A 多参数，多返回值
func A(a int, b string) (x int, y string) {
	return 1, ""
}

//B 多参数，多返回值，类型合并
func B(a, b int) (x, y int) {
	return 1, 2
}

//C 返回值，不命名
func C() (int, int) {
	return 1, 2
}

//D 返回值，命名
func D() (x, y int) {
	a, b := 1, 2
	return a, b
}

//E 不定长变参
func E(a ...interface{}) {
	fmt.Println(a)
}

//F 匿名函数
func F() {
	func() {
		fmt.Println("Hello")
	}()
}

func closure(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func defer1() {
	fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")
}

func defer2() {
	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}
}

func defer3() {
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}

func p1() {
	fmt.Println("p1")
}

func p2() {
	fmt.Println("p2")
}

func panic1() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover work")
		}
	}()
	panic("panic occur")
}

func test() {
	var fs = [4]func(){}

	for i := 0; i < 4; i++ {
		defer fmt.Println("defer i=", i)
		defer func() { fmt.Println("defer_closure i=", i) }()
		fs[i] = func() { fmt.Println("closure i=", i) }
	}

	for _, f := range fs {
		f()
	}
}
