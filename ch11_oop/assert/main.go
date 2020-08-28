package main

import "fmt"

type point struct {
	x int
	y int
}

func main() {
	var a interface{}
	var pp point = point{1, 2}
	a = pp
	var b point
	b = a.(point)
	fmt.Println(b)

	var x interface{}
	var bb float64 = 1.1
	x = bb
	if y, ok := x.(float32); ok {
		fmt.Printf("y的類型是%T,值是%v\n", y, y)
	} else {
		fmt.Println("轉換失敗")
	}
	if z, ok := x.(float64); ok {
		fmt.Printf("y的類型是%T,值是%v\n", z, z)
	} else {
		fmt.Println("轉換失敗")
	}

}
