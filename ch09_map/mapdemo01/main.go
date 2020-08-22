package main

import "fmt"

func main() {
	var a map[string]string
	// 在使用map前，需要先make，make的作用是分配記憶體空間
	a = make(map[string]string)
	a["no1"] = "tom"
	a["no2"] = "cat"
	a["no1"] = "acg"
	a["no3"] = "tom"
	fmt.Println(a)
}
