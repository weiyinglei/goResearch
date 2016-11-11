package main

import "fmt"

//A A
type A struct {
	Name string
}

//B B
type B struct {
	A
	Name string
}

type M struct{}

func testTen() {
	b := B{Name: "B's Name", A: A{Name: "A's Name"}}
	fmt.Println(b)
	fmt.Println(b.Name)
	fmt.Println(b.A.Name)
}

/*
 * 声明，使用
 * 类型别名与方法
 * method value method expression
 * 方法名称冲突与字段访问权限
 */
func main() {
	// testTen()
	valueAndExpression()
}

func valueAndExpression() {
	m := M{}
	m.method1()
	(*M).method1(&m)
}

func (m *M) method1() {
	fmt.Println("A")
}
