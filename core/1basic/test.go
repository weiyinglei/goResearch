/*
 * test.go
 */
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var t int64
	t = 100000000000
	Numcpus := runtime.NumCPU()
	runtime.GOMAXPROCS(Numcpus)
	Nowtime := time.Now()
	c := make(chan int64, Numcpus)
	for i := 0; i < Numcpus; i++ {
		go add(t*int64(i)/int64(Numcpus), t*int64(i+1)/int64(Numcpus), c)
	}
	var s int64

	for i := 0; i < Numcpus; i++ {
		s += <-c
	}
	Endtime := time.Now()
	fmt.Println("逻辑CPU个数", Numcpus, "\n\r总数是", s, "\n\r用时", Endtime.Sub(Nowtime))
}
func add(s, e int64, c chan int64) {
	var t int64
	t = 0
	for i := s; i < e; i++ {
		t += i
	}
	c <- t
}
