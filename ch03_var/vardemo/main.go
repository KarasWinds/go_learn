package main

import "fmt"

// 定義全局變數
var gn = 100
var gk = "roy"
var ge = 10.1
var (
	gn2 = 23
	gk2 = "ben"
	ge2 = 12.3
)

func main() {
	// 定義變數
	var i int
	// 設定變數
	// i = 10
	// 使用變數
	fmt.Println("i=", i)
	// 自行偵測變數型態
	var num = 10.1
	fmt.Println("num=", num)
	// 直接聲明變數值，變數應為初次使用
	name := "tom"
	fmt.Println("name=", name)
	// 設定多重變數
	var n1, n2, n3 int
	fmt.Println("n1:", n1, "n2:", n2, "n3:", n3)
	var n, k, r = 100, "you", 10.2
	fmt.Println("n:", n, "k:", k, "r:", r)
	nn, kk, rr := 1, "re", 4.5
	fmt.Println("nn:", nn, "kk:", kk, "rr:", rr)
	fmt.Println("gn:", gn, "gk:", gk, "ge:", ge)
	fmt.Println("gn2:", gn2, "gk2:", gk2, "ge2:", ge2)
}
