package main

import "fmt"

func main() {
	var arr [5]int = [...]int{34, 56, 321, 76, 432}
	var arrSlice = arr[1:4]
	for i := 0; i < len(arrSlice); i++ {
		fmt.Printf("arrSlice[%v]:%v\n", i, arrSlice[i])
	}
	fmt.Println()
	for i, v := range arrSlice {
		fmt.Printf("arrSlice[%v]:%v\n", i, v)
	}

	var intSlice []int = []int{100, 200, 300}
	intSlice = append(intSlice, 400, 500, 600)
	fmt.Println("intSlice:", intSlice)
	intSlice = append(intSlice, intSlice...)
	fmt.Println("intSlice:", intSlice)
}
