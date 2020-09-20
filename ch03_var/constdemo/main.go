package main

import "fmt"

func main() {
	const (
		x = 9
		y = 0
	)

	fmt.Println(x)
	fmt.Println(y)

	const (
		a = iota
		b
		c, d = iota, iota
	)

	fmt.Println(a, b, c, d)
}
