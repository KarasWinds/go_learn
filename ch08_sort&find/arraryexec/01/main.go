package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var intArr [10]int

	for i := 0; i < len(intArr); i++ {
		rand.Seed(time.Now().UnixNano())
		intArr[i] = rand.Intn(100) + 1
	}
	fmt.Println(intArr)

	sum := 0.0
	var maxNum int = intArr[0]
	var minNum int = intArr[0]
	var num int = 55
	var maxIndex int
	var minIndex int
	var numIndex int = -1
	for i := 0; i < len(intArr); i++ {
		sum += float64(intArr[i])
		if intArr[i] == num {
			numIndex = i
		}
		if intArr[i] > maxNum {
			maxIndex = i
			maxNum = intArr[i]
		} else if intArr[i] < minNum {
			minIndex = i
			minNum = intArr[i]
		}
		fmt.Print(" ", intArr[len(intArr)-i-1])
	}
	fmt.Println()

	fmt.Println("平均值 :", sum/float64(len(intArr)))
	fmt.Printf("最大值%v,位置在%v\n", maxNum, maxIndex)
	fmt.Printf("最小值%v,位置在%v\n", minNum, minIndex)

	if numIndex == -1 {
		fmt.Printf("陣列中沒有%v\n", num)
	} else {
		fmt.Printf("陣列中有%v,位置在%v\n", num, numIndex)
	}
}
