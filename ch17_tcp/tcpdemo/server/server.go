package main

import (
	"fmt"
	"io"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	// fmt.Printf("Server Watting for client, IP:%s\n", conn.RemoteAddr().String())
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err == io.EOF {
			fmt.Printf("Client:%v exit\n", conn.RemoteAddr().String())
			return
		} else if err != nil {
			fmt.Println("server Read err:", err)
			return
		}
		fmt.Print(string(buf[:n]))
	}
}

func main() {

	fmt.Println("server start listening...")
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}
	defer listen.Close()

	for {
		fmt.Println("waitting listen")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err:", err)
		} else {
			fmt.Printf("Accept() suc con:%v, client IP:%v\n", conn, conn.RemoteAddr().String())
		}
		go process(conn)
	}

	// fmt.Printf("listen suc:%v\n", listen)
}
