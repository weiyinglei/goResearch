/*
 * defineThread.go
 */
package main

import (
	"fmt"
	"runtime"
	"time"
)

func SayHello() {
	for i := 0; i < 10; i++ {
		fmt.Print("Hello")
		runtime.Gosched()
	}
}

func SayWorld() {
	for i := 0; i < 10; i++ {
		fmt.Println("World")
		runtime.Gosched()
	}
}

func main() {
	go SayHello()
	go SayWorld()
	time.Sleep(5 * time.Second)
}
