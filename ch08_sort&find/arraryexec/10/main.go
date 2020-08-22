package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	var intArr [8]int
	var sum int
	for i := 0; i < len(intArr); i++ {
		rand.Seed(time.Now().UnixNano())
		intArr[i] = rand.Intn(100)
		sum += intArr[i]
	}
	fmt.Println(intArr)

	maxScore := intArr[0]
	minScore := intArr[0]
	var maxIndex int
	var minIndex int
	for i := 1; i < len(intArr); i++ {
		if intArr[i] > maxScore {
			maxScore = intArr[i]
			maxIndex = i
		} else if intArr[i] < minScore {
			minScore = intArr[i]
			minIndex = i
		}
	}

	fmt.Printf("最高分:第%v評委，分數%v\n", maxIndex+1, maxScore)
	fmt.Printf("最低分:第%v評委，分數%v\n", minIndex+1, minScore)

	avg := float64(sum-maxScore-minScore) / float64(len(intArr)-2)
	fmt.Printf("扣除最高最低分之平均分數:%.2f\n", avg)

	var absScore float64
	maxIndex, minIndex = 0, 0
	maxAbs, minAbs := math.Abs(avg-float64(intArr[0])), math.Abs(avg-float64(intArr[0]))
	for i := 1; i < len(intArr); i++ {
		absScore = math.Abs(avg - float64(intArr[i]))
		if absScore > maxAbs {
			maxAbs = absScore
			maxIndex = i
		} else if absScore < minAbs {
			minAbs = absScore
			minIndex = i
		}
	}
	fmt.Printf("最佳評委:第%v位，分數%v\n", minIndex+1, intArr[minIndex])
	fmt.Printf("最差評委:第%v位，分數%v\n", maxIndex+1, intArr[maxIndex])
}
