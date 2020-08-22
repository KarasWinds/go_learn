package main

import "fmt"

func main() {
	var price float32 = 22.2
	fmt.Println("price=", price)
	var num3 float32 = -123.0000901
	var num4 float64 = -123.0000901
	fmt.Println("num3=", num3)
	fmt.Println("num4=", num4)
	// 浮點數預設為float64
	var num5 = 1.2
	fmt.Printf("浮點數預設為%T\n", num5)
	// 科學符號表示
	num6 := 1.23454e3
	num7 := 1.23454E3
	num8 := 1.23454E-3
	fmt.Println("num6:", num6, "num7:", num7, "num8:", num8)
}
