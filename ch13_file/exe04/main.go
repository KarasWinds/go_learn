package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	filePath := "./abc.txt"
	// os.O_APPEND 在黨案內容後追加
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open file err:%v\n", err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		fmt.Print(str)
		if err == io.EOF {
			break
		}
	}

	str := "hello, TAIPEI\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 1; i++ {
		writer.WriteString(str)
	}
	// writer 是先寫入暫存區的，需要再寫入檔案
	writer.Flush()

}
