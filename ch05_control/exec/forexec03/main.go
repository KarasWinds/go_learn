package main

import (
	"fmt"
)

func main() {
	var totalLevel int = 3

	for i := 1; i <= totalLevel; i++ {
		for j := 1; j <= totalLevel-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			if k == 1 || k == 2*i-1 || i == totalLevel {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	for i := 1; i <= 2*totalLevel-1; i++ {
		if i > totalLevel {
			for j := 1; j <= i-totalLevel; j++ {
				fmt.Print(" ")
			}
			for k := 1; k <= 2*i-1; k++ {
				if k == 1 || k == 2*i-1-(i-totalLevel)*4 {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
			}
		} else {
			for j := 1; j <= totalLevel-i; j++ {
				fmt.Print(" ")
			}
			for k := 1; k <= 2*i-1; k++ {
				if k == 1 || k == 2*i-1 {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
			}
		}
		fmt.Println()
	}

	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			fmt.Printf("%v*%v=%v\t", i, j, i*j)
		}
		fmt.Println()
	}
}
