package main

import "fmt"

type person struct {
	Name string
	Age  int
}

func main() {
	p2 := person{"ann", 18}
	fmt.Println(p2)

	var p3 *person = new(person)
	(*p3).Name = "tom"
	p3.Age = 26
	fmt.Println(*p3)

	var p4 *person = &person{}
	(*p4).Name = "kai"
	p4.Age = 30
	fmt.Println(*p4)
}
