package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var intArr [5]int
	for i := 0; i < len(intArr); i++ {
		rand.Seed(time.Now().UnixNano())
		intArr[i] = rand.Intn(10)
	}

	var maxNum int = intArr[0]
	var minNum int = intArr[0]
	var maxIndex int
	var minIndex int
	for i := 1; i < len(intArr); i++ {
		if intArr[i] > maxNum {
			maxNum = intArr[i]
			maxIndex = i
		} else if intArr[i] < minNum {
			minNum = intArr[i]
			minIndex = i
		}
	}
	fmt.Println(intArr)
	fmt.Printf("最大數位置在%v為%v\n", maxIndex, maxNum)
	fmt.Printf("最小數位置在%v為%v\n", minIndex, minNum)

}
