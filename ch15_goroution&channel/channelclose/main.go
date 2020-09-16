package main

import "fmt"

func main() {
	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200

	// 關閉後只能讀取
	close(intChan)

	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan2 <- i * 2
	}
	// 關閉後使用for range到結尾才不出錯
	close(intChan2)
	for v := range intChan2 {
		fmt.Println("v:", v)
	}
}
