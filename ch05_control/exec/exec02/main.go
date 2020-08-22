package main

import (
	"fmt"
	"math"
)

func main() {
	var score int
	fmt.Println("輸入成績:")
	fmt.Scanln(&score)
	if score == 100 {
		fmt.Println("獎品BMW")
	} else if score >= 80 && score < 100 {
		fmt.Println("獎品ip7p")

	} else if score >= 60 && score < 80 {
		fmt.Println("獎品ipad")
	} else {
		fmt.Println("沒有獎勵")
	}

	var bt bool = true
	if bt == false {
		fmt.Println("a")
	} else if bt {
		fmt.Println("b")
	} else if !bt {
		fmt.Println("c")
	} else {
		fmt.Println("d")
	}

	var a float64 = 2.0
	var b float64 = 4.0
	var c float64 = 2.0

	m := b*b - 4*a*c
	if m > 0 {
		fmt.Println((-b + math.Sqrt(m)) / 2 / a)
		fmt.Println((-b - math.Sqrt(m)) / 2 / a)
	} else if m == 0 {
		fmt.Println((-b - math.Sqrt(m)) / 2 / a)
	} else {
		fmt.Print("無解")
	}

}
