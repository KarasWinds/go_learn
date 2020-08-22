package main

import "fmt"

type integer int

func (i integer) print() {
	fmt.Println(i)
}

func (i *integer) change() {
	*i++
}

type student struct {
	Name string
	Age  int
}

func (stu *student) String() string {
	return fmt.Sprintf("Name:%v,Age:%v", stu.Name, stu.Age)
}

func main() {
	var i integer = 10
	i.print()
	i.change()
	i.print()

	stu1 := student{
		Name: "ann",
		Age:  28,
	}
	fmt.Println(&stu1)
	fmt.Printf("%v\n", &stu1)
}
