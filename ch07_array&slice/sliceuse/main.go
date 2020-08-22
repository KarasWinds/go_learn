package main

import "fmt"

func main() {
	var varSlice []float64 = make([]float64, 5, 10)
	varSlice[1] = 10
	varSlice[3] = 20
	fmt.Println("varSlice:", varSlice)
	fmt.Println("varSlice的size:", len(varSlice))
	fmt.Println("varSlice的cap:", cap(varSlice))

	var intSlice []int = []int{3, 6, 22}
	fmt.Println("intSlice:", intSlice)
	fmt.Println("varSlice的size:", len(intSlice))
	fmt.Println("varSlice的cap:", cap(intSlice))
}
