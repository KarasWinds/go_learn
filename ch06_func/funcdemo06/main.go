package main

import "fmt"

var (
	fun1 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

func main() {
	res1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)

	fmt.Println("res1 =", res1)

	a := func(n1 int, n2 int) int {
		return n1 - n2
	}
	res2 := a(20, 10)
	fmt.Println("res2 =", res2)
	res3 := a(40, 10)
	fmt.Println("res3 =", res3)
	res4 := fun1(20, 3)
	fmt.Println("res4 =", res4)

}
