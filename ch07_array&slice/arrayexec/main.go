package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var myChar [26]byte
	for i := 0; i < 26; i++ {
		myChar[i] = byte('A' + i)
	}
	for _, value := range myChar {
		fmt.Printf("%c ", value)
	}
	fmt.Println()

	var intArr = [...]int{1, 2, 54, -5, -45, 45}
	maxVal := intArr[0]
	maxValIndex := 0

	for index, value := range intArr {
		if value > maxVal {
			maxVal, maxValIndex = value, index
		}
	}

	fmt.Printf("maxVal:%v,maxValIndex:%v\n", maxVal, maxValIndex)

	var intArr2 = [...]int{1, 2, 54, -5, -45, 45}
	sum := 0
	for _, value := range intArr2 {
		sum += value
	}
	fmt.Printf("sum:%v,avg:%.2v\n", sum, float64(sum)/float64(len(intArr2)))

	var intArr3 [5]int
	var arrLen int = len(intArr3)
	for i := 0; i < arrLen; i++ {
		rand.Seed(time.Now().UnixNano())
		intArr3[i] = rand.Intn(100)
	}
	fmt.Println(intArr3)

	for i := 0; i < arrLen/2; i++ {
		temp := intArr3[arrLen-1-i]
		intArr3[arrLen-1-i] = intArr3[i]
		intArr3[i] = temp
	}
	fmt.Println(intArr3)
}
