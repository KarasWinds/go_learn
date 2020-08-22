package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var b = false
	fmt.Println("b=", b)
	fmt.Println(unsafe.Sizeof(b))
}
