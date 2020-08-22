package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%v是奇數\n", i)
	}

	var positiveCount int
	var negativeCount int
	var num int

	for {
		fmt.Println("請輸入一個整數，0為結束")
		fmt.Scanln(&num)
		if num > 0 {
			positiveCount++
			continue
		} else if num < 0 {
			negativeCount++
			continue
		} else {
			fmt.Printf("輸入的正數有%v個，負數有%v個\n", positiveCount, negativeCount)
			break
		}
	}

}
