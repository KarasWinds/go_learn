package main

import "fmt"

func main() {
	var arr = [2][3]int{{1, 2, 3}, {4, 5, 6}}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("%v\t", arr[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
	for _, arrVal := range arr {
		for _, v := range arrVal {
			fmt.Printf("%v\t", v)
		}
		fmt.Println()
	}
}
