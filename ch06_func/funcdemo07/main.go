package main

import (
	"fmt"
	"strings"
)

func addUpper() func(int) int {
	var n int = 10
	return func(x int) int {
		n += x
		return n
	}
}

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			name += suffix
		}
		return name
	}
}

func main() {
	f := addUpper()
	fmt.Printf("%T\n", addUpper())
	fmt.Println(f(1))
	fmt.Println(f(3))

	s := makeSuffix(".jpg")
	fmt.Println("檔名處理後:", s("aaa"))
	fmt.Println("檔名處理後:", s("bbb.jpg"))
}
