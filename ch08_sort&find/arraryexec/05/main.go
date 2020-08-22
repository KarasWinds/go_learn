package main

import "fmt"

func main() {
	var intArr [5]int = [5]int{1, 3, 5, 7, 9}
	fmt.Println(intArr)
	for i := 0; i < len(intArr); i++ {
		fmt.Print(" ", intArr[len(intArr)-i-1])
	}
	fmt.Println()
}
