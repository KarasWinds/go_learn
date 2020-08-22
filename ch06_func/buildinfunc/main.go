package main

import "fmt"

func main() {
	num1 := 100
	fmt.Printf("num1的類型:%T、值:%v、位址:%v\n", num1, num1, &num1)

	num2 := new(int)
	*num2 = 200
	fmt.Printf("num2的類型:%T、值:%v、位址:%v、指向的值:%v\n", num2, num2, &num2, *num2)
}
