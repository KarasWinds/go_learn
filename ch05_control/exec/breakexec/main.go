package main

import "fmt"

func main() {
	var sum int
	for i := 1; i <= 100; i++ {
		sum += i
		if sum > 20 {
			fmt.Println("當總和大於20,i =", i)
			break
		}
	}

	var name string
	var pwd string
	var loginChance = 3
	for i := 1; i <= loginChance; i++ {
		fmt.Println("請輸入姓名:")
		fmt.Scanln(&name)
		fmt.Println("請輸入密碼:")
		fmt.Scanln(&pwd)
		if name == "Tom" && pwd == "1234" {
			fmt.Println("登入成功")
			break
		}
		if i == 3 {
			fmt.Println("登入失敗，請重來")
			break
		}
		fmt.Printf("登入失敗，還有%v機會\n", loginChance-i)

	}
}
