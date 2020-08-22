package main

import "fmt"

func main() {
	var age byte
	fmt.Println("請輸入您的年齡:")
	fmt.Scanln(&age)
	if age >= 18 {
		fmt.Println("您已經成年，請為自己行為負責")
	} else {
		fmt.Println("您還未成年喔!!")
	}
}
