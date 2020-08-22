package main

import "fmt"

func main() {
	var second float64
	var gender string
	fmt.Println("請輸入秒數:")
	fmt.Scanln(&second)
	if second <= 8 {
		fmt.Println("請輸入性別:")
		fmt.Scanln(&gender)
		if gender == "M" {
			fmt.Println("進入男子組決賽")
		} else {
			fmt.Println("進入女子組決賽")
		}
	}

	var month byte
	var age byte
	fmt.Println("請輸入遊玩月份:")
	fmt.Scanln(&month)
	fmt.Println("請輸入遊客年齡:")
	fmt.Scanln(&age)
	if month >= 4 && month <= 10 {
		if age > 60 {
			fmt.Println("敬老票，售價20")
		} else if age < 18 {
			fmt.Println("兒童票，售價30")
		} else {
			fmt.Println("全票，售價60")
		}
	} else {
		if age < 60 && age > 18 {
			fmt.Println("全票，售價40")
		} else {
			fmt.Println("半票，售價20")
		}

	}

}
