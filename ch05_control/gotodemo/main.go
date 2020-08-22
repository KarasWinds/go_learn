package main

import "fmt"

func main() {
	fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")
	if true {
		goto label1
	}
	fmt.Println("4")
	fmt.Println("5")
label1:
	fmt.Println("6")
	fmt.Println("7")
	fmt.Println("8")
	fmt.Println("9")
}
