package main

import "fmt"

func main() {
	var key byte
	fmt.Println("請輸入一個字元:(a,b,c,d,e,f,g)")
	fmt.Scanf("%c", &key)

	switch key {
	case 'a':
		fmt.Println("周一")
	case 'b':
		fmt.Println("周二")
	case 'c':
		fmt.Println("周三")
	case 'd':
		fmt.Println("周四")
	case 'e':
		fmt.Println("周五")
	case 'f':
		fmt.Println("周六")
	case 'g':
		fmt.Println("周日")
	default:
		fmt.Println("輸入錯誤")
	}

	var num int = 10
	switch num {
	case 10:
		fmt.Println("ok1")
		fallthrough
	case 20:
		fmt.Println("ok2")
	case 30:
		fmt.Println("ok3")
	default:
		fmt.Println("NO~")
	}
}
