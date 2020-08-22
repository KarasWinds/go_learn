package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var intArr [10]int = [10]int{0, 14, 26, 36, 42, 57, 63, 77, 81, 99}
	fmt.Println(intArr)

	rand.Seed(time.Now().UnixNano())
	var num int = rand.Intn(100) + 1
	fmt.Println("插入數字:", num)

	var intArr2 [11]int

	for i := 0; i < len(intArr2); i++ {
		if intArr[i] >= num {
			intArr2[i] = num
			for j := i + 1; j < len(intArr2); j++ {
				intArr2[j] = intArr[j-1]
			}
			break
		}
		intArr2[i] = intArr[i]
	}

	fmt.Println(intArr2)

}
