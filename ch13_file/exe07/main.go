package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type charCount struct {
	chCount    int
	numCount   int
	spaceCount int
	otherCount int
}

func main() {
	fileName := "./abc.txt"
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Printf("open file err:%v", err)
		return
	}
	defer file.Close()

	var count charCount

	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n')
		for _, v := range str {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				count.chCount++
			case v == ' ' || v == '\t':
				count.spaceCount++
			case v >= '0' && v <= '9':
				count.numCount++
			default:
				count.otherCount++

			}
		}
		if err == io.EOF {
			break
		}
	}
	fmt.Printf("英文數:%v個,數字:%v個,空格%v個,其他%v個\n", count.chCount, count.numCount, count.spaceCount, count.otherCount)
}
