package main

import "fmt"

// Cat is a class
type Cat struct {
	Name string
	Age  int
}

func main() {
	allChan := make(chan interface{}, 3)

	allChan <- 10
	allChan <- "cat"
	cat := Cat{"tom", 4}
	allChan <- cat
	<-allChan
	<-allChan
	newCat := <-allChan

	fmt.Printf("chanCat類型%T,值%v\n", newCat, newCat)

	c := newCat.(Cat)
	fmt.Printf("newCat.Name:%v\n", c.Name)
}
