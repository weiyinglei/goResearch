package main

import (
	"fmt"
)

/*
 * define
 * reslice
 * append
 * share
 * copy
 */
func main() {
	// define1()
	// define2()
	// define3()
	// define4()

	// append1()
	// share()
	// copy1()
	test()
}

func test() {
	a := [...]int{1, 2, 3, 4, 5}
	fmt.Println(a, len(a), cap(a))
	s := a[0:len(a)]
	fmt.Println(s, len(s), cap(s))
}

func define1() {
	var s1 []int
	fmt.Println(s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
}

func define2() {
	a := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a)
	s1 := a[5:10]
	fmt.Println(s1)

	s2 := a[5:len(a)]
	fmt.Println(s2)

	s3 := a[5:]
	fmt.Println(s3)

	s4 := a[:5]
	fmt.Println(s4)
}

func define3() {
	s1 := make([]int, 3, 10)
	fmt.Println(s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
}

func define4() {
	a := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k'}
	fmt.Println(len(a), cap(a))

	s1 := a[3:5]
	fmt.Println(len(s1), cap(s1))
	fmt.Println(s1)
}

func append1() {
	s1 := make([]int, 3, 6)
	fmt.Printf("%p %v", s1, s1)
	s1 = append(s1, 1, 2, 3)
	fmt.Printf("\n\r%p %v", s1, s1)
	s1 = append(s1, 4, 5, 6)
	fmt.Printf("\n\r%p %v", s1, s1)
}

func share() {
	a := [...]int{1, 2, 3, 4, 5}
	s1 := a[2:5]
	s2 := a[1:3]
	fmt.Println(s1, s2)
	s2 = append(s2, 1, 1, 1, 1, 1, 1, 1, 1)
	s1[0] = 10
	fmt.Println(s1, s2)
}

func copy1() {
	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := []int{7, 8, 9}
	// copy(s1, s2)
	// fmt.Println(s1)
	copy(s2, s1)
	fmt.Println(s2)
}
