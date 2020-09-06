package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filePath := "./abc.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err:%v\n", err)
		return
	}

	defer file.Close()

	str := "hello, Gardon\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	// writer 是先寫入暫存區的，需要再寫入檔案
	writer.Flush()

}
