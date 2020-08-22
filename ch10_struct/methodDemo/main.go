package main

import "fmt"

type person struct {
	Name string
}

func (a person) test() {
	a.Name = "kai"
	fmt.Println("test()", a.Name)
}

func (a person) speak() {
	fmt.Println("Hi", a.Name)
}

func (a person) cal() {
	res := 0
	for i := 0; i <= 1000; i++ {
		res += i
	}
	fmt.Println(a.Name, ":,計算結果=", res)
}

func (a person) cal2(n int) {
	res := 0
	for i := 0; i <= n; i++ {
		res += i
	}
	fmt.Println(a.Name, ":", res)
}

func (a person) getSum(n1 int, n2 int) int {
	return n1 + n2
}

func main() {
	var p person
	p.Name = "ann"
	p.test()
	fmt.Println("main()", p.Name)
	p.speak()
	p.cal()
	p.cal2(5)
	n1 := 10
	n2 := 20
	res := p.getSum(n1, n2)
	fmt.Println(res)
}
