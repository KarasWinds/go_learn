package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file := "./test"
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("read file error :%v\n", err)
	}

	fmt.Printf("%v\n", string(content))
}
