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
	for i, j := 0, len(str); i < j; i++ {
		if string(str[i]) == " " {
			continue
		}
		fmt.Println(string(str[i]))
	}

outer:
	for i := 0; i < 5; i++ {
		for k := 0; k < 5; k++ {
			if i > 0 {
				break outer
			}
			fmt.Println(k)
		}
	}
}
