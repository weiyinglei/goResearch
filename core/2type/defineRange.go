/*
 * defineRange.go
 */
package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}
	sli := arr[0:3]
	mp := map[int]string{1: "a", 2: "b", 3: "c"}
	for k, v := range mp {
		fmt.Println(k, "=", v)
	}
	for k, v := range arr {
		fmt.Println(k, "=", v)
	}
	for k, v := range sli {
		fmt.Println(k, "=", v)
	}
	for k, v := range arr {
		fmt.Println(k, ":", v, ":", mp[v])
	}
}
