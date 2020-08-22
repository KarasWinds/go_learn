package main

import "fmt"

func main() {
	str := "hello@World"
	slice := str[6:]
	fmt.Println("slice :", slice)

	arr1 := []byte(str)
	arr1[0] = 'z'
	str = string(arr1)
	fmt.Println("str :", str)
	arr2 := []rune(str)
	arr2[0] = '台'
	str = string(arr2)
	fmt.Println("str :", str)
}
