package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("參數有", len(os.Args))
	for i, v := range os.Args {
		fmt.Printf("第%v參數為%v\n", i, v)
	}
}
