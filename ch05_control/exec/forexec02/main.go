package main

import "fmt"

func main() {
	var classNum int = 3
	var stuNum int = 5
	var passCount int
	total := 0.0
	for j := 1; j <= classNum; j++ {
		sum := 0.0
		for i := 1; i <= stuNum; i++ {
			var score float64
			fmt.Printf("請輸入第%d班、第%d個學生成績\n", j, i)
			fmt.Scanln(&score)
			sum += score
			if score >= 60 {
				passCount++
			}
		}
		fmt.Printf("第%d班的平均分是%v\n", j, sum/float64(stuNum))
		total += sum
	}
	fmt.Printf("全部班的平均分是%v\n", total/float64(stuNum)/float64(classNum))
	fmt.Printf("及格人數為%d\n", passCount)

}
