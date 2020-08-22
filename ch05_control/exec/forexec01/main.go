package main

import "fmt"

func main() {
	var max int = 100
	var count int
	var sum int
	for i := 1; i <= max; i++ {
		if i%9 == 0 {
			count++
			sum += i
			fmt.Printf("第%v個，為%v\n", count, i)
		}
	}
	fmt.Printf("總共%v個，總和%v\n", count, sum)

	var num int = 6
	for i := 0; i <= num; i++ {
		fmt.Printf("%v + %v = %v\n", i, num-i, num)
	}
}
