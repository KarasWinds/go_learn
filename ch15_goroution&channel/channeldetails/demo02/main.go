package main

import (
	"fmt"
	"time"
)

func main() {

	intChan := make(chan int, 10)
	for i := 1; i <= 10; i++ {
		intChan <- i
	}

	stringChan := make(chan string, 5)
	for i := 1; i <= 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}

	for {
		select {
		case v := <-intChan:
			fmt.Printf("intChan:%d\n", v)
			time.Sleep(time.Second)
		case v := <-stringChan:
			fmt.Printf("stringChan:%s\n", v)
			time.Sleep(time.Second)
		default:
			fmt.Println("讀取結束")
			return
		}
	}
}
