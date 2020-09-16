package main

import "fmt"

func main() {
	var intChan chan int
	intChan = make(chan int, 3)

	fmt.Printf("intChan:%v\n", intChan)

	intChan <- 10
	num := 200
	intChan <- num
	fmt.Printf("intChan長度%v,容量%v\n", len(intChan), cap(intChan))

	num2 := <-intChan
	fmt.Printf("intChan長度%v,容量%v\n", len(intChan), cap(intChan))
	fmt.Println(num2)
}
