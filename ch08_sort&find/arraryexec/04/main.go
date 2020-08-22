package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var intArr [4][4]int
	for i := 0; i < len(intArr); i++ {
		for j := 0; j < len(intArr[i]); j++ {
			rand.Seed(time.Now().UnixNano())
			intArr[i][j] = rand.Intn(10)
			fmt.Print(intArr[i][j], " ")
		}
		fmt.Println()
	}

	fmt.Println()

	tempArr := intArr[0]
	intArr[0] = intArr[3]
	intArr[3] = tempArr
	tempArr = intArr[1]
	intArr[1] = intArr[2]
	intArr[2] = tempArr

	for i := 0; i < len(intArr); i++ {
		for j := 0; j < len(intArr[i]); j++ {
			fmt.Print(intArr[i][j], " ")
		}
		fmt.Println()
	}

}
