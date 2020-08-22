package main

import "fmt"

func modifyUser(user map[string]map[string]string, name string) {
	if user[name] != nil {
		user[name]["pwd"] = "888888"
	} else {
		fmt.Println("新增使用者")
		fmt.Println("請輸入使用者匿名:")
		var nickname string
		fmt.Scanln(&nickname)
		fmt.Println("請輸入使用者密碼")
		var pwd string
		fmt.Scanln(&pwd)
		user[name] = make(map[string]string)
		user[name]["nickname"] = nickname
		user[name]["pwd"] = pwd
	}
}

func main() {
	user := make(map[string]map[string]string)
	user["ann"] = make(map[string]string)
	user["ann"]["nickname"] = "ian"
	user["ann"]["pwd"] = "password"
	fmt.Println(user)

	fmt.Println("請輸入使用者名稱:")
	var name string
	fmt.Scanln(&name)
	modifyUser(user, name)

	fmt.Println(user)
}
