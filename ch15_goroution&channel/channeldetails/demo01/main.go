package main

import "fmt"

func main() {
	var chan1 chan int
	chan1 = make(chan int)
	fmt.Println(chan1)

	// 唯寫Channel
	var chan2 chan<- int
	chan2 = make(chan int)
	fmt.Println(chan2)

	// 唯讀Channel
	var chan3 <-chan int
	chan3 = make(chan int)
	fmt.Println(chan3)
}
