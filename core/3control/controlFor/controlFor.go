/*
 * controlFor.go
 * for init;condition;post{}
 * for condition{}
 * for {}
 */
package main

import "fmt"

func main() {
	str := "hello world"
	for i := 0; i < len(str); i++ {
		fmt.Print(string(str[i]))
	}
	fmt.Println()

	outer:
	for i := 0; i < 5; i++ {
		for k := 0; k < 5; k++ {
			if i > 0 {
				break outer
			}
			fmt.Print(k)
		}
	}
}
