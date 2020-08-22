package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var intArr [8]int
	var sum int
	for i := 0; i < len(intArr); i++ {
		rand.Seed(time.Now().UnixNano())
		intArr[i] = rand.Intn(10)
		sum += intArr[i]
	}

	avg := float64(sum) / float64(len(intArr))
	fmt.Println(intArr)
	fmt.Printf("平均為%v\n", avg)

	var bigcount int
	var smallcount int
	for i := 0; i < len(intArr); i++ {
		if float64(intArr[i]) > avg {
			bigcount++
		} else if float64(intArr[i]) < avg {
			smallcount++
		}
	}
	fmt.Printf("大於平均有%v個數\n", bigcount)
	fmt.Printf("小於平均有%v個數\n", smallcount)
}
