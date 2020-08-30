package main

import (
	"fmt"
	"go_learn/ch12_project/familyAccount/utils"
)

func main() {

	fmt.Println("透過物件導向方式實作")
	utils.NewFamilyAccount().MainMenu()
}
