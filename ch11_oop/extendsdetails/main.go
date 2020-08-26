package main

import "fmt"

// AA is a class
type AA struct {
	Name string
	age  int
}

func (a *AA) sayOk() {
	fmt.Println("AA sayOk", a.Name)
}

func (a *AA) hello() {
	fmt.Println("AA hello", a.Name)
}

// BB is a class
type BB struct {
	AA
	Name string
}

func (b *BB) sayOk() {
	fmt.Println("BB sayOk", b.Name)
}

func main() {
	var b BB
	b.Name = "aki"
	b.age = 18
	b.sayOk()
	b.hello()

	b.AA.Name = "bin"
	b.sayOk()
	b.hello()

}
