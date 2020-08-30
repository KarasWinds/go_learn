package main

import "fmt"

func main() {
	// 接收輸入控制
	var key string
	// 控制退出迴圈
	var loop bool = true
	// 帳戶餘額
	var blance float64 = 10000.0
	// 收支費用
	var money float64
	// 收支說明
	var note string
	// 收支明細字串
	var details string = "----------------------收支明細----------------------\n"
	details += "收支\t\t金額\t\t餘額\t\t說明\n"
	var detailsFlag bool = false

	// 顯示主選單
	for {
		key = ""
		fmt.Println("--------------------家庭收支軟體--------------------")
		fmt.Println("                    1. 收支明細")
		fmt.Println("                    2. 收入登記")
		fmt.Println("                    3. 支出登記")
		fmt.Println("                    4.   退出")
		fmt.Print("請選擇1-4:")
		fmt.Scanln(&key)
		switch key {
		case "1":
			if detailsFlag {
				fmt.Println(details)
			} else {
				fmt.Println("尚無收支明細")
			}
		case "2":
			fmt.Print("本次收入金額:")
			fmt.Scanln(&money)
			blance += money
			fmt.Print("收入說明:")
			fmt.Scanln(&note)
			details += fmt.Sprintf("收入\t\t%v\t\t%v\t\t%v\n", money, blance, note)
			detailsFlag = true
		case "3":
			fmt.Print("本次支出金額:")
			fmt.Scanln(&money)
			if money > blance {
				fmt.Println("餘額不足")
				break
			}
			blance -= money
			fmt.Print("支出說明:")
			fmt.Scanln(&note)
			details += fmt.Sprintf("支出\t\t%v\t\t%v\t\t%v\n", money, blance, note)
			detailsFlag = true
		case "4":
			var keyYesNo string
			for {
				fmt.Print("確定要退出嗎?y/n")
				fmt.Scanln(&keyYesNo)
				if keyYesNo == "y" {
					loop = false
					break
				} else if keyYesNo == "n" {
					break
				} else {
					fmt.Println("輸入錯誤，請重新輸入")
				}
			}
		default:
			fmt.Println("請輸入正確選項")
		}
		if !loop {
			break
		}
	}
}
