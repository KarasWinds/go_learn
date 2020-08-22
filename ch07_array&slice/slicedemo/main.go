package main

import "fmt"

func main() {
	var intArr [5]int = [...]int{1, 2, 3, 4, 5}
	// slice引用1開始~不包含3
	intSlice := intArr[1:3]
	fmt.Println("intArr:", intArr)
	fmt.Printf("intArr的位置:%p\n", &intArr)
	fmt.Printf("intArr[1]的位置:%p\n", &intArr[1])
	fmt.Println("intSlice:", intSlice)
	fmt.Printf("intSlice[0]:%v\n", &intSlice[0])
	fmt.Printf("intSlice的位置:%p\n", &intSlice)
	fmt.Println("intSlice的個數:", len(intSlice))
	fmt.Println("intSlice的容量:", cap(intSlice))
}
