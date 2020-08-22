package main

import "fmt"

func main() {
	var c1 byte = 'a'
	var c2 byte = '0'
	fmt.Println("c1=", c1)
	fmt.Println("c2=", c2)
	fmt.Printf("c1=%c c2=%c\n", c1, c2)
	var c3 int = '北'
	fmt.Println("c3=", c3)
	fmt.Printf("c3=%c\n", c3)
	var cc = '南'
	fmt.Printf("cc=%T\n", cc)
	var c4 int = 25105
	fmt.Printf("c4=%c\n", c4)
	var n1 = 10 + 'a'
	fmt.Println("n1=", n1)
}
