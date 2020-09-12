package main

import "fmt"

func addUpper(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	return res
}
func main() {
	res := addUpper(10)
	if res == 55 {
		fmt.Printf("addUpper正確,返回值:%v,預期值:%v\n", res, 55)
	} else {
		fmt.Printf("addUpper錯誤,返回值:%v,預期值:%v\n", res, 55)
	}
}
