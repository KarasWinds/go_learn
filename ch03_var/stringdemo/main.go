package main

import "fmt"

func main() {
	var address string = "Taipei"
	fmt.Println("address:", address)
	str1 := "abc\nabc\n"
	str2 := `abc\nabc\n`
	fmt.Println(str1, str2)
	var str = "hello" + "world"
	str += "!!"
	fmt.Println(str)
}
