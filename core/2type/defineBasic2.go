/*
 * defineBasic.go
 * 简单类型
 * bool            1    false
 * byte            1    0      uint8
 * rune            4    0      int32
 * int,uint        4|8  0      32|64
 * int8,uint8      1    0      -128~127,0~255
 * int16,uint16    2    0      -32768~32767,0~65535
 * int32,uint32    4    0      -21亿~21亿,0~42亿
 * int64,uint64    8    0
 * float32         4    0.0
 * float64         8    0.0
 * complex64       8
 * complex128      16
 * string
 *
 * 复杂类型
 * array
 * slice           nil
 * map             nil
 * uintptr         4|8
 * function        nil
 * struct
 * interface       nil
 *
 */
package main

import (
	"fmt"
	"reflect"
)

type data struct {
	a int
}

func main() {
	d := data{1234}
	var p *data

	p = &d
	fmt.Println(p)
	fmt.Println(reflect.ValueOf(p))
	fmt.Println(reflect.ValueOf(d))
}
