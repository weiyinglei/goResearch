package main

import (
	"fmt"
)

type person struct {
	Name    string
	Age     int
	Contact struct {
		Phone, City string
	}
}

type human struct {
	Sex  int
	Name string
}

type teacher struct {
	human
	person
	// Name string
	Age int
}
type student struct {
	human
	person
	// Name string
	Age int
}

/*
 * 定义，使用
 * 字面值初始化
 * 匿名结构
 * 匿名字段
 * 结构间的赋值与比较
 * 嵌入结构
 */
func main() {
	// define1()
	// anonymity1()
	// anonymity2()
	// anonymity3()
	embed()
}

func embed() {
	t := teacher{
		human: human{
			Name: "Kate",
		},
		Age: 18,
	}

	t.human.Name = "Kate for human"
	t.person.Name = "Kate for person"

	fmt.Println(t)
}

//匿名字段
// func anonymity3() {
// 	s := student{
// 		"Kate",
// 		18,
// 	}
// 	fmt.Println(s)
// }

//匿名结构
func anonymity2() {
	p := &person{
		Name: "Kate",
		Age:  18,
	}
	p.Contact.Phone = "18888888888"
	p.Contact.City = "杭州"
	fmt.Println(p)
}

//匿名结构
func anonymity1() {
	p := struct {
		Name string
		Age  int
	}{
		Name: "Kate",
		Age:  18,
	}
	fmt.Println(p)
}

//定义，使用
func define1() {
	p := person{
		Name: "Tom",
		Age:  20,
	}

	fmt.Println(p)
	a(p)
	b(&p)
	fmt.Println(p)
}

func a(p person) {
	p.Age = 15
	fmt.Println(p)
}

func b(p *person) {
	p.Age = 25
	fmt.Println(p)
}
