package main

import "fmt"

func printpyramid(totalLevel int) {
	for i := 1; i <= totalLevel; i++ {
		for j := 1; j <= totalLevel-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func printmulti(num int) {
	for i := 1; i <= num; i++ {
		for j := 1; j <= num; j++ {
			fmt.Printf("%v*%v=%v\t", i, j, i*j)
		}
		fmt.Println()
	}
}

func main() {
	var n int
	fmt.Println("請輸入金字塔層數")
	fmt.Scanln(&n)
	printpyramid(n)

	fmt.Println("請輸入乘法表層數")
	fmt.Scanln(&n)
	printmulti(n)
}
