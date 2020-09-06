package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file1 := "./abc.txt"
	file2 := "./def.txt"

	content, err := ioutil.ReadFile(file1)
	if err != nil {
		fmt.Printf("open file err:%v", err)
		return
	}

	err = ioutil.WriteFile(file2, content, 0666)
	if err != nil {
		fmt.Printf("write file err:%v", err)
		return
	}
}
