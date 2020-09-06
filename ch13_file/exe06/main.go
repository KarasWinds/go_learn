package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func copyFile(dstFileName string, srcFileNmae string) (written int64, err error) {
	srcFile, err := os.OpenFile(srcFileNmae, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Printf("open file err:%v", err)
		return
	}
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)

	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err:%v", err)
		return
	}
	defer dstFile.Close()
	writer := bufio.NewWriter(dstFile)

	return io.Copy(writer, reader)
}

func main() {
	srcFile := "./images.jpg"
	dstFile := "./copy.jpg"
	_, err := copyFile(dstFile, srcFile)
	if err == nil {
		fmt.Printf("copy file sucess\n")
	} else {
		fmt.Printf("cpoy file err:%v\n", err)
	}
}
