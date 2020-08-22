package main

import "fmt"

func swap(n1 *int, n2 *int) {
	t := *n1
	*n1 = *n2
	*n2 = t
}

func main() {
	a := 10
	b := 20
	fmt.Printf("a=%v,b=%v\n", a, b)
	swap(&a, &b)
	fmt.Printf("a=%v,b=%v\n", a, b)
}
