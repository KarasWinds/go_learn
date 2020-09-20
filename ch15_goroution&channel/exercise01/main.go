package main

import "fmt"

var (
	num    int = 200
	thread int = 8
	count  int
)

func writeData(numChan chan int) {
	for i := 1; i <= num; i++ {
		numChan <- i
	}
	close(numChan)
}

func readData(numChan chan int, resChan chan map[int]int) {
	for {
		resMap := make(map[int]int)
		i, ok := <-numChan
		if !ok {
			count++
			break
		}
		res := calRes(i)
		resMap[i] = res
		fmt.Println(resMap)
		resChan <- resMap
	}
}

func calRes(n int) int {
	if n == 0 {
		return 0
	}
	return n + calRes(n-1)
}

func main() {
	numChan := make(chan int, num)
	resChan := make(chan map[int]int, num)

	go writeData(numChan)

	for i := 1; i <= thread; i++ {
		go readData(numChan, resChan)
	}

	for {
		if count == thread {
			close(resChan)
			break
		}
	}

	for i := 1; i <= num; i++ {
		v, ok := <-resChan
		if !ok {
			break
		}
		fmt.Printf("res[%v]:%v\n", i, v)
	}

}
