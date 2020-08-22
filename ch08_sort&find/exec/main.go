package main

import "fmt"

func main() {
	var scores [3][5]float64
	for i := 0; i < len(scores); i++ {
		for j := 0; j < len(scores[i]); j++ {
			fmt.Printf("請輸入第%v第%v個學生成績:", i, j)
			fmt.Scanln(&scores[i][j])
		}
	}

	totalSum := 0.0
	for i, arr := range scores {
		sum := 0.0
		for _, v := range arr {
			sum += v
		}
		totalSum += sum
		fmt.Printf("第%v班總分%v平均%v分\n", i+1, sum, sum/float64(len(scores[i])))
	}

	fmt.Printf("所有班級總分為%v平均分為%v\n", totalSum, totalSum/15)
}
