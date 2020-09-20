package main

import "fmt"

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	for {
		n, ok := <-intChan
		if !ok {
			exitChan <- true
			break
		}
		for i := 2; i <= n; i++ {
			if n == i {
				primeChan <- n
				break
			}
			if n%i == 0 {
				break
			}
		}
	}
}

func main() {

	num := 10000
	thread := 6
	intChan := make(chan int, num)
	primeChan := make(chan int)
	exitChan := make(chan bool, thread)

	go func() {
		for i := 1; i <= num; i++ {
			intChan <- i
		}
		close(intChan)
	}()

	for i := 1; i <= thread; i++ {
		go primeNum(intChan, primeChan, exitChan)
		// fmt.Printf("開啟線程%v\n", i)
	}

	go func() {
		for i := 1; i <= thread; i++ {
			<-exitChan
			// fmt.Printf("結束線程%v\n", i)
		}
		close(primeChan)
	}()

	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Printf("質數:%d\n", res)
	}
	fmt.Println("退出main()")
}
