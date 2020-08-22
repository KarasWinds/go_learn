package main

import "fmt"

func bubbleSort(arr *[5]int) {
	var temp int
	for i := 0; i < len(*arr)-1; i++ {
		for j := 0; j < len(*arr)-i-1; j++ {
			temp = (*arr)[j]
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
			}
		}
	}
}

func main() {
	arr := [5]int{24, 69, 80, 57, 13}
	fmt.Println("排序前:", arr)
	bubbleSort(&arr)
	fmt.Println("排序後:", arr)
}
