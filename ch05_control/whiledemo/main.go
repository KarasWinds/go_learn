package main

import "fmt"

func main() {
	var i int = 1
	for {
		if i > 10 {
			break
		}
		fmt.Println("Hello World!", i)
		i++
	}

	var j int = 1
	for {
		fmt.Println("Hello !", j)
		j++
		if j > 10 {
			break
		}
	}
}
