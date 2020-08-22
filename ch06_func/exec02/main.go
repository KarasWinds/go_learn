package main

import "fmt"

func fc(n int) int {
	if n == 1 {
		return 3
	}
	return 2*fc(n-1) + 1
}

func main() {
	n := 2
	fmt.Printf("f(%v)=%v\n", n, fc(n))
}
