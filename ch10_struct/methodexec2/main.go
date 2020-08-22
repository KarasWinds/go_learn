package main

import "fmt"

type methodUtils struct {
}

func (mu methodUtils) print() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 8; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (mu methodUtils) print2(m int, n int) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (mu methodUtils) area(l float64, w float64) float64 {
	return l * w
}

func main() {
	var mu methodUtils
	mu.print()
	fmt.Println()
	mu.print2(6, 4)

	fmt.Println(mu.area(3.5, 4.6))
}
