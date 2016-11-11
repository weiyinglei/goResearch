/*
 * controlSwitch.go
 */
package main

import "fmt"

func main() {
	i := 5
	switch i {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3, 4, 5, 6:
		fmt.Println("3,4,5,6")
		fallthrough
	default:
		fmt.Println("others")
	}
}
