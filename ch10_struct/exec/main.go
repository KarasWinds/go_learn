package main

import "fmt"

type person struct {
	Name string
	Age  int
}

func main() {
	var p1 person = person{"ann", 28}
	var p2 *person = &p1
	fmt.Println(p1)
	fmt.Println(p2)
	p2.Age = 32
	fmt.Printf("p1的記憶體位址:%p,值為%v\n", &p1, p1)
	fmt.Printf("p2的記憶體位址:%p,值為%p\n", &p2, p2)
}
