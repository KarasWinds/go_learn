package main

import "fmt"

func test(slice []int) {
	slice[0] = 100
}

func main() {
	var slice1 []int
	var arr []int = []int{1, 2, 3, 4, 5}
	slice1 = arr[:]
	var slice2 = slice1
	slice2[0] = 10

	fmt.Println("slice2:", slice2)
	fmt.Println("slice1:", slice1)
	fmt.Println("arr:", arr)
	test(slice1)
	fmt.Println("slice2:", slice2)
	fmt.Println("slice1:", slice1)
	fmt.Println("arr:", arr)
}
