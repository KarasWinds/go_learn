package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var count int
	for {
		rand.Seed(time.Now().UnixNano())
		num := rand.Intn(100) + 1
		fmt.Println("num =", num)
		count++
		if num == 99 {
			break
		}
	}
	fmt.Printf("總共花了幾%v次\n", count)

lable0: //自定義標籤
	for i := 0; i <= 4; i++ {
		for j := 0; j <= 5; j++ {
			if j == 2 {
				break lable0
			}
			fmt.Println("j =", j)
		}
	}
}
