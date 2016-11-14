/*
 * defineMap.go
 */
package main

import "fmt"

func main() {
	map1 := make(map[string]string) /*key是字符串类型，值也是字符串类型*/
	map1["a"] = "1"
	map1["b"] = "2"
	map1["pi"] = "3.1415926"
	map1["sh"] = "上海"
	v, ok := map1["sh"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("key 'sh'不存在")
	}
	v, ok = map1["bj"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("key 'bj'不存在")
	}
}
