package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filePath := "./abc.txt"
	// os.O_TRUNC 先清空檔案內容
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("open file err:%v\n", err)
		return
	}

	defer file.Close()

	str := "hello, Go\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}
	// writer 是先寫入暫存區的，需要再寫入檔案
	writer.Flush()

}
