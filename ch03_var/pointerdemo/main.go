package main

import "fmt"

func main() {
	var i int = 10
	fmt.Printf("i的值:%v\n", i)
	fmt.Printf("i的位置:%v\n", &i)

	var ptr *int = &i
	fmt.Printf("ptr的值:%v\n", ptr)
	fmt.Printf("ptr的位置:%v\n", &ptr)
	fmt.Printf("ptr指向的值:%v\n", *ptr)
	*ptr = 15
	fmt.Printf("i的值:%v\n", i)

	var a int = 300
	var b int = 400
	var p *int = &a
	*p = 100
	p = &b
	*p = 200
	fmt.Printf("a的值:%v,b的值:%v\n", a, b)

}
