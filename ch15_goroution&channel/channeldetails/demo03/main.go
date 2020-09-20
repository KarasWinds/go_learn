package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Printf("hello world %d\n", i)
	}
}

func test() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("test() err :", err)
		}
	}()

	var myMap map[int]string
	myMap[1] = "test"
}

func main() {

	go sayHello()
	go test()

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Printf("main() %d\n", i)
	}
}
