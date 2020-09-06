package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./test")
	if err != nil {
		fmt.Println("open file err=", err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		fmt.Print(str)
		if err == io.EOF { //io.EOF is end of file
			break
		}
	}
	fmt.Println("文件讀取結束")
}
