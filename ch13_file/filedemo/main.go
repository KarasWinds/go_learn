package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./test")
	if err != nil {
		fmt.Println("open file err=", err)
	}
	fmt.Printf("file=%v\n", file)
	err = file.Close()
	if err != nil {
		fmt.Println("close file err=", err)
	}
}
