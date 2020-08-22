package main

import "fmt"

func main() {
	var intArr [3]int
	fmt.Println(intArr)
	fmt.Printf("add:%p\n", &intArr)
	fmt.Printf("add[0]:%p\n", &intArr[0])
	fmt.Printf("add[1]:%p\n", &intArr[1])

	// var score [5]float64
	// for i := 0; i < len(score); i++ {
	// 	fmt.Printf("請輸入第%d個值:\n", i+1)
	// 	fmt.Scanln(&score[i])
	// }
	// for i := 0; i < len(score); i++ {
	// 	fmt.Printf("score[%d]:%v\n", i, score[i])
	// }

	var numArr01 [3]int = [3]int{1, 2, 3}
	fmt.Println("numArr01:", numArr01)
	var numArr02 = [3]int{1, 2, 3}
	fmt.Println("numArr02:", numArr02)
	var numArr03 = [...]int{1, 2, 3, 4}
	fmt.Println("numArr03:", numArr03)
	var numArr04 = [...]int{1: 800, 0: 342, 2: 432, 4: 543}
	fmt.Println("numArr04:", numArr04)
	numArr05 := [...]string{1: "tom", 0: "cat", 2: "369", 3: "kz"}
	fmt.Println("numArr05:", numArr05)

}
