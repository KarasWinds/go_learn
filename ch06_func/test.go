package main

import "fmt"

var name string = "mouse"

func test1() {
	name := "cat"
	fmt.Println("name =", name)
}

func test2() {
	name = "cat"
	fmt.Println("name =", name)
}

func main() {

	fmt.Println("name =", name)
	test1()
	fmt.Println("name =", name)
	test2()
	fmt.Println("name =", name)
}
