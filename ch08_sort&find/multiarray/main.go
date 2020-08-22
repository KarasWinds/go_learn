package main

import "fmt"

func main() {
	var arr [4][6]int
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println()

	arr[1][2] = 1
	arr[2][1] = 2
	arr[2][3] = 3

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}

	var arr2 [2][3]int

	fmt.Printf("arr2[0]的位置%p\n", &arr2[0])
	fmt.Printf("arr2[1]的位置%p\n", &arr2[1])
	fmt.Printf("arr2[0][0]的位置%p\n", &arr2[0][0])
	fmt.Printf("arr2[1][0]的位置%p\n", &arr2[1][0])
	fmt.Println()

}
