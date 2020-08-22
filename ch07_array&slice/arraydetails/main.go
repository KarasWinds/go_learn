package main

import "fmt"

func test01(arr [3]int) {
	arr[0] = 88
}

func test02(arr *[3]int) {
	(*arr)[0] = 88
}

func main() {
	var arr01 [3]int
	var arr02 [3]string
	var arr03 [3]bool
	fmt.Printf("arr01:%v\narr02:%v\narr03:%v\n", arr01, arr02, arr03)

	arr := [...]int{22, 33, 44}
	test01(arr)
	fmt.Println(arr)
	test02(&arr)
	fmt.Println(arr)
}
