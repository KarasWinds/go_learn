package main

import "fmt"

func main() {
	var i int = 5
	fmt.Printf("%b\n", i)
	i = 012
	fmt.Println("i=", i)
	i = 0x23
	fmt.Println("i=", i)

}
