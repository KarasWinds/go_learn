package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i uint = 1
	fmt.Println("i=", i)
	fmt.Printf("資料型態:%T\n", i)
	fmt.Printf("占用byte數:%d\n", unsafe.Sizeof(i))
}
