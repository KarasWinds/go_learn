package main

import "fmt"

// Account is a class struct
type Account struct {
	AccountNo string
	Pwd       string
	Balance   float64
}

func (account *Account) deposite(money float64, pwd string) {
	if pwd != account.Pwd {
		fmt.Println("密碼不正確")
		return
	}
	if money <= 0 {
		fmt.Println("金額不正確")
		return
	}
	account.Balance += money
	fmt.Println("存款成功,餘額:", account.Balance)
}

func (account *Account) withDraw(money float64, pwd string) {
	if pwd != account.Pwd {
		fmt.Println("密碼不正確")
		return
	}
	if money <= 0 {
		fmt.Println("金額不正確")
		return
	} else if money > account.Balance {
		fmt.Println("餘額不足")
		return
	}
	account.Balance -= money
	fmt.Println("提款成功,餘額:", account.Balance)
}

func (account *Account) query(pwd string) {
	if pwd != account.Pwd {
		fmt.Println("密碼不正確")
		return
	}
	fmt.Println("餘額:", account.Balance)
}

func main() {
	var user = Account{
		AccountNo: "369",
		Pwd:       "999",
		Balance:   0,
	}

	user.deposite(0, "333")
	user.deposite(-1, "999")
	user.deposite(1000, "999")

	user.withDraw(0, "333")
	user.withDraw(-1, "999")
	user.withDraw(5000, "999")
	user.withDraw(500, "999")

	user.query("333")
	user.query("999")

}
