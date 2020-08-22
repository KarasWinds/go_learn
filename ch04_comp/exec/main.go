package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20
	fmt.Printf("交換前a=%d,b=%d\n", a, b)
	a += b
	b = a - b
	a = a - b
	fmt.Printf("交換後a=%d,b=%d\n", a, b)

	var n1 int = 30
	var n2 int = 40
	var max int
	if n1 > n2 {
		max = n1
	} else {
		max = n2
	}
	fmt.Println("max=", max)

	var n3 int = 50
	if n3 > max {
		max = n3
	}
	fmt.Println("max=", max)
}
