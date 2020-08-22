package main

import (
	"errors"
	"fmt"
)

func test1() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("error =", err)
			// 進行錯誤處理
		}
	}()

	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res =", res)
}

func readConf(name string) (err error) {
	if name == "config.ini" {
		return nil
	}
	return errors.New("讀取文件錯誤")
}

func test2() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("error =", err)
			// 進行錯誤處理
		}
	}()
	err := readConf("config.ii")
	if err != nil {
		panic(err)
	}
	fmt.Println("test2()...")
}

func main() {
	test1()
	test2()
	fmt.Println("main()...")

}
