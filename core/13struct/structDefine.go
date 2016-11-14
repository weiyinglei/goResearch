/*
 * structDefine.go
 * 结构体
 */
package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	class string
}

func (this Student) getName() string {
	return this.Name
}

func (this *Student) getAge() int {
	return this.Age
}

func (this Student) SetName1(name string) {
	this.Name = name
}

func (this *Student) SetName(name string) {
	this.Name = name
}

func main() {
	s1 := new(Student)
	s1.Name = "张三"
	s1.Age = 10
	s1.class = "4班"
	fmt.Println(s1)
	s2 := Student{"李四", 11, "8班"}
	s3 := Student{Name: "王五", Age: 12, class: "2班"}
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s1.getName(), s1.getAge())

	s2.SetName1("张一")
	fmt.Println(s2)
	s2.SetName("张二")
	fmt.Println(s2)
}
