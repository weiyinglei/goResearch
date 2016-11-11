package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println(time.Now())
	var s string
	for i := 0; i < len(os.Args); i++ {
		s += os.Args[i] + "\n"
	}
	fmt.Println(s)
}
