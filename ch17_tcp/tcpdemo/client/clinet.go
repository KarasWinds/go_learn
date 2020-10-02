package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("client dial err:", err)
		return
	}
	reader := bufio.NewReader(os.Stdin)
	for {

		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("ReadString err:", err)
		}
		line = strings.Trim(line, "\r\n")
		if line == "exit" {
			fmt.Println("connect exit")
			break
		}
		_, err = conn.Write([]byte(line + "\n"))
		if err != nil {
			fmt.Println("conn.Write err:", err)
		}
	}
}
