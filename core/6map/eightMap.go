package main

import (
	"fmt"
)

/*
 * define
 * use
 * n demension
 * iteration
 * slice[map]
 */
func main() {
	// define1()
	// use()
	// iteration()
	test()
}

func test() {
	m1 := map[int]string{1: "a", 2: "b", 3: "c"}
	fmt.Println(m1)
	m2 := make(map[string]int)
	for k, v := range m1 {
		m2[v] = k
	}
	fmt.Println(m2)
}

func define1() {
	var m1 map[int]string
	fmt.Println(m1)
	var m2 = map[int]string{}
	fmt.Println(m2)
	var m3 = make(map[int]string)
	fmt.Println(m3)
	m4 := make(map[int]string)
	fmt.Println(m4)
}

func use() {
	m := make(map[int]string)
	m[1] = "OK"
	fmt.Println(m)
	fmt.Println(m[1])
	delete(m, 1)
	fmt.Println(m)
}

func iteration() {
	s := []int{1, 2, 3, 4, 5, 8, 6}
	for i, v := range s {
		fmt.Println(i, v)
	}
	m := make(map[int]int)
	m[0] = 1
	m[1] = 2
	m[2] = 3
	for k, v := range m {
		fmt.Println(k, v)
	}
}
