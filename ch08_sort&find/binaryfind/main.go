package main

import "fmt"

func binaryFind(arr *[6]int, leftIndex int, rightIndex int, findVal int) {
	if leftIndex > rightIndex {
		fmt.Println("找不到", findVal)
		return
	}
	middleIndex := (leftIndex + rightIndex) / 2
	if findVal > (*arr)[middleIndex] {
		binaryFind(arr, middleIndex+1, rightIndex, findVal)
	} else if findVal < (*arr)[middleIndex] {
		binaryFind(arr, leftIndex, middleIndex-1, findVal)
	} else {
		fmt.Printf("找到了%v，在位置%v\n", findVal, middleIndex)
	}
}

func main() {
	arr := [6]int{1, 8, 89, 100, 1234}
	binaryFind(&arr, 0, len(arr), -10)

}
