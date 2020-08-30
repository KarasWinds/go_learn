package utils

import "fmt"

// FamilyAccount is a class
type FamilyAccount struct {
	key string
	// 控制退出迴圈
	loop bool
	// 帳戶餘額
	blance float64
	// 收支費用
	money float64
	// 收支說明
	note string
	// 收支明細字串
	details string
	// 收支確認
	detailsFlag bool
}

// MainMenu 主選單
func (fA *FamilyAccount) MainMenu() {
	for {
		fA.key = ""
		fmt.Println("--------------------家庭收支軟體--------------------")
		fmt.Println("                    1. 收支明細")
		fmt.Println("                    2. 收入登記")
		fmt.Println("                    3. 支出登記")
		fmt.Println("                    4.   退出")
		fmt.Print("請選擇1-4:")
		fmt.Scanln(&fA.key)
		switch fA.key {
		case "1":
			fA.showDetail()
		case "2":
			fA.income()
		case "3":
			fA.pay()
		case "4":
			fA.exit()
		default:
			fmt.Println("請輸入正確選項")
		}
		if !fA.loop {
			break
		}
	}
}

func (fA *FamilyAccount) showDetail() {
	if fA.detailsFlag {
		fmt.Println(fA.details)
	} else {
		fmt.Println("尚無收支明細")
	}
}

func (fA *FamilyAccount) income() {
	fmt.Print("本次收入金額:")
	fmt.Scanln(&fA.money)
	fA.blance += fA.money
	fmt.Print("收入說明:")
	fmt.Scanln(&fA.note)
	fA.details += fmt.Sprintf("收入\t\t%v\t\t%v\t\t%v\n", fA.money, fA.blance, fA.note)
	fA.detailsFlag = true
}

func (fA *FamilyAccount) pay() {
	fmt.Print("本次支出金額:")
	fmt.Scanln(&fA.money)
	if fA.money > fA.blance {
		fmt.Println("餘額不足")
	}
	fA.blance -= fA.money
	fmt.Print("支出說明:")
	fmt.Scanln(&fA.note)
	fA.details += fmt.Sprintf("支出\t\t%v\t\t%v\t\t%v\n", fA.money, fA.blance, fA.note)
	fA.detailsFlag = true
}

func (fA *FamilyAccount) exit() {
	var keyYesNo string
	for {
		fmt.Print("確定要退出嗎?y/n")
		fmt.Scanln(&keyYesNo)
		if keyYesNo == "y" {
			fA.loop = false
			break
		} else if keyYesNo == "n" {
			break
		} else {
			fmt.Println("輸入錯誤，請重新輸入")
		}
	}
}

// NewfamilyAccount 建構子
func NewfamilyAccount() *FamilyAccount {

	return &FamilyAccount{
		loop:        true,
		blance:      10000.0,
		details:     "----------------------收支明細----------------------\n收支\t\t金額\t\t餘額\t\t說明\n",
		detailsFlag: false,
	}

}
