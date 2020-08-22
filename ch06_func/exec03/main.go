package main

import "fmt"

func exec(l int, n int) int {
	if n > 1 {
		n--
		l := 2 * (l + 1)
		return exec(l, n)
	}
	return l
}

func main() {
	n := 10
	l := 1
	fmt.Printf("第%v天剩%v，最初有%v\n", n, l, exec(l, n))
}
