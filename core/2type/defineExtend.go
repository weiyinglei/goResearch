/*
 * defineExtend.go
 */
package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	class string
}

func (this *Student) Display() {
	fmt.Println(this.Name, ",", this.Age)
}

type CollageStudent struct {
	Student
	Profession string
}

func (this *CollageStudent) Display() {
	fmt.Println(this.Name, ",", this.Profession)
}

func main() {
	s1 := CollageStudent{Student: Student{Name: "李四", Age: 23, class: "二班"}, Profession: "物理"}
	s1.Display()
	fmt.Println(s1.Student.Name)
	fmt.Println(s1.Student.Age)
	fmt.Println(s1.Name)
	fmt.Println(s1.Age)
	fmt.Println(s1.Profession)
}
