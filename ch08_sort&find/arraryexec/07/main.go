package main

import (
	"fmt"
	"math/rand"
	"time"
)

func binaryFind(arr *[10]int, num int, numIndex *int, minIndex int, maxIndex int) {
	if minIndex > maxIndex {
		return
	}
	middleIndex := (minIndex + maxIndex) / 2
	if num < (*arr)[middleIndex] {
		binaryFind(arr, num, numIndex, minIndex, middleIndex-1)
	} else if num > (*arr)[middleIndex] {
		binaryFind(arr, num, numIndex, middleIndex+1, maxIndex)
	} else {
		*numIndex = middleIndex
		return
	}
}

func main() {
	var intArr [10]int
	for i := 0; i < len(intArr); i++ {
		rand.Seed(time.Now().UnixNano())
		intArr[i] = rand.Intn(100) + 1
	}
	fmt.Println(intArr)

	var temp int
	for i := 0; i < len(intArr); i++ {
		for j := 0; j < len(intArr)-1-i; j++ {
			if intArr[j] > intArr[j+1] {
				temp = intArr[j]
				intArr[j] = intArr[j+1]
				intArr[j+1] = temp
			}
		}
	}
	fmt.Println(intArr)

	num := 90
	numIndex := -1

	binaryFind(&intArr, num, &numIndex, 0, len(intArr)-1)
	if numIndex == -1 {
		fmt.Printf("找不到%v\n", num)
	} else {
		fmt.Printf("在%v位置找到%v\n", numIndex, num)
	}

}
