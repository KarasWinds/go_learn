package main

import "fmt"

func main() {
	var a int32 = 23
	var b int32 = 32
	if a+b > 50 {
		fmt.Println("Hello World")
	}

	var c float64 = 30.2
	var d float64 = 2.3
	if c > 10.0 && d < 20.0 {
		fmt.Println(c + d)
	}

	var num1 int32 = 3
	var num2 int32 = 12
	num := num1 + num2

	if num%3 == 0 && num%5 == 0 {
		fmt.Println(num, "能被3也能被5整除")
	}
	var year int = 2020
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		fmt.Println(year, "是閏年")
	}
}
