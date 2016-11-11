/*
 * controlIfElse.go
 */
package main

import "fmt"

func main() {
	if a := 2; a < 2 {
		fmt.Println("a<2")
	} else {
		fmt.Println("a=", a)
	}
}
