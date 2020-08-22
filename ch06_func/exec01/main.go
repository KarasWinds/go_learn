package main

import "fmt"

func fbn(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fbn(n-1) + fbn(n-2)
}

func main() {
	n := 6
	fmt.Printf("第%v個費波納數為:%v\n", n, fbn(n))
}
