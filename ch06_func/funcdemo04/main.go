package main

import "fmt"

func test1(n int) {
	n += 10
	fmt.Println("test1() n =", n)
}
func test2(n *int) {
	*n += 10
	fmt.Println("test2() n =", *n)
}

func main() {
	n := 10
	test1(n)
	fmt.Println("main() n =", n)
	test2(&n)
	fmt.Println("main() n =", n)
}
