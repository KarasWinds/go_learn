package model

import "fmt"

type account struct {
	name    string
	pwd     string
	balance float64
}

func NewAccount(name string, pwd string, balance float64) *account {
	if len(name) < 6 && len(name) > 10 {
		return nil
	}
	if len(pwd) != 6 {
		return nil
	}
	if balance < 1000 {
		return nil
	}
	return &account{
		name:    name,
		pwd:     pwd,
		balance: balance,
	}
}

func (account *account) SetName(name string) {
	if len(name) < 6 && len(name) > 10 {
		fmt.Println("名稱長度不對")
	} else {
		account.name = name
	}
}

func (account *account) SetPwd(pwd string) {
	if len(pwd) != 6 {
		fmt.Println("密碼長度不對")
	} else {
		account.pwd = pwd
	}
}

func (account *account) Deposite(money float64, pwd string) {
	if pwd != account.pwd {
		fmt.Println("密碼不正確")
		return
	}
	if money <= 0 {
		fmt.Println("金額不正確")
		return
	}
	account.balance += money
	fmt.Println("存款成功,餘額:", account.balance)
}

func (account *account) WithDraw(money float64, pwd string) {
	if pwd != account.pwd {
		fmt.Println("密碼不正確")
		return
	}
	if money <= 0 {
		fmt.Println("金額不正確")
		return
	} else if money > account.balance {
		fmt.Println("餘額不足")
		return
	}
	account.balance -= money
	fmt.Println("提款成功,餘額:", account.balance)
}

func (account *account) Query(pwd string) {
	if pwd != account.pwd {
		fmt.Println("密碼不正確")
		return
	}
	fmt.Println("餘額:", account.balance)
}
