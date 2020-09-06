package main

import (
	"flag"
	"fmt"
)

func main() {

	// 定義參數，用於接收命令列參數
	var user string
	var pwd string
	var host string
	var port int

	flag.StringVar(&user, "u", "", "用戶名(預設為空)")
	flag.StringVar(&pwd, "pwd", "", "密碼(預設為空)")
	flag.StringVar(&host, "h", "localhost", "主機名(預設為localhost)")
	flag.IntVar(&port, "p", 3306, "連接port(預設為3306)")

	// 轉換(必要)
	flag.Parse()

	fmt.Printf("用戶名:%v,密碼:%v,主機名:%v,port:%v\n", user, pwd, host, port)
}
