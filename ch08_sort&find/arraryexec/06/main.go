package main

import "fmt"

func main() {
	stringArr := [5]string{"AA", "BB", "CC", "DD", "AA"}

	for i, v := range stringArr {
		if v == "AA" {
			fmt.Printf("找到AA,在位置%v\n", i)
		}
	}
}
