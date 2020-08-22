package main

import "fmt"

func main() {
	fmt.Println(2 & 3)
	fmt.Println(2 | 3)
	fmt.Println(2 ^ 3)
	fmt.Println(-2 ^ 2)

	a := 1 >> 2
	b := 1 << 2
	fmt.Println("a =", a, "b =", b)
}
