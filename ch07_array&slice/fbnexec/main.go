package main

import "fmt"

func fbn(num int) []uint64 {
	var fbnSlice = make([]uint64, num)
	fbnSlice[0] = 1
	fbnSlice[1] = 1
	for i := 2; i < num; i++ {
		fbnSlice[i] = fbnSlice[i-1] + fbnSlice[i-2]
	}
	return fbnSlice
}

func main() {
	fbnSlice := fbn(20)
	fmt.Println("fnbSlice :", fbnSlice)
}
