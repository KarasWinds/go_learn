package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var intArr [3][4]int
	for i := 0; i < len(intArr); i++ {
		for j := 0; j < len(intArr[i]); j++ {
			rand.Seed(time.Now().UnixNano())
			intArr[i][j] = rand.Intn(10)
			fmt.Print(intArr[i][j], " ")
		}
		fmt.Println()
	}

	fmt.Println()

	for i := 0; i < len(intArr); i++ {
		for j := 0; j < len(intArr[i]); j++ {
			if i == 0 || i == len(intArr)-1 || j == 0 || j == len(intArr[i])-1 {
				intArr[i][j] = 0
			}
			fmt.Print(intArr[i][j], " ")
		}
		fmt.Println()
	}

}
