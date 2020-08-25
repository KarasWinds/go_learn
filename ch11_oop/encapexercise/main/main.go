package main

import (
	"fmt"
	"go_learn/ch11_oop/encapexercise/model"
)

func main() {
	acc := model.NewAccount("369369", "999999", 1000)
	fmt.Println(acc)
	acc.SetName("knight")
	acc.SetPwd("000000")
	fmt.Println(acc)
}
