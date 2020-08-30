package main

import (
	"fmt"
	"go_learn/ch12_project/customerManager/model"
	"go_learn/ch12_project/customerManager/service"
)

type customerView struct {
	// 用戶輸入key
	key string
	// 循環控制
	loop bool

	customerService *service.CustomerService
}

func (cv *customerView) list() {
	customers := cv.customerService.CustomerList()
	if len(customers) == 0 {
		fmt.Println("----------------尚無客戶資料----------------")
		return
	}
	fmt.Println("----------------客戶列表開始----------------")
	fmt.Println("ID\t姓名\t性別\t年齡\t電話\tE-mail")
	for _, v := range customers {
		fmt.Println(v.GetInfo())
	}
	fmt.Println("----------------客戶列表結束----------------")
}

func (cv *customerView) add() {
	var (
		name   string
		gender string
		age    int
		phone  string
		email  string
	)
	fmt.Println("----------------開始建立客戶----------------")
	fmt.Printf("姓名: ")
	fmt.Scanln(&name)
	fmt.Printf("性別: ")
	fmt.Scanln(&gender)
	fmt.Printf("年齡: ")
	fmt.Scanln(&age)
	fmt.Printf("電話: ")
	fmt.Scanln(&phone)
	fmt.Printf("E-mail: ")
	fmt.Scanln(&email)
	fmt.Println("----------------結束建立客戶----------------")
	var newCustomer model.Customer = model.NewCustomer(name, gender, age, phone, email)
	cv.customerService.Add(newCustomer)
	fmt.Println("客戶建立完成")
}

func (cv *customerView) delete() {
	fmt.Println("------------------刪除客戶------------------")
	for {
		fmt.Println("請輸入待刪除客戶編號(-1退出):")
		var id int
		fmt.Scanln(&id)
		if id == -1 {
			break
		} else if id == 0 {
			continue
		}
		choice := ""
		fmt.Printf("確定是否刪除ID:%v客戶(Y/N):", id)
		fmt.Scanln(&choice)
		if choice == "Y" || choice == "y" {
			if cv.customerService.Delete(id) {
				fmt.Printf("已刪除ID:%v客戶\n", id)
				break
			} else {
				fmt.Printf("無此ID:%v客戶\n", id)
			}
		}
	}
}

func (cv *customerView) edit() {
	var (
		name   string
		gender string
		age    int
		phone  string
		email  string
	)
	fmt.Println("------------------修改客戶------------------")
	for {
		fmt.Println("請輸入待修改客戶編號(-1退出):")
		var id int
		fmt.Scanln(&id)
		if id == -1 {
			break
		} else if id == 0 {
			continue
		}
		if cv.customerService.Check(id) {
			coustomer := cv.customerService.Get(id)
			fmt.Printf("姓名(%v):", coustomer.GetName())
			fmt.Scanln(&name)
			fmt.Printf("性別(%v):", coustomer.GetGender())
			fmt.Scanln(&gender)
			fmt.Printf("年齡(%v):", coustomer.GetAge())
			fmt.Scanln(&age)
			fmt.Printf("phone(%v):", coustomer.GetPhone())
			fmt.Scanln(&phone)
			fmt.Printf("E-mail(%v):", coustomer.GetEmail())
			fmt.Scanln(&email)
			cv.customerService.Update(id, name, gender, age, phone, email)
			fmt.Println("------------------修改完成------------------")
			break
		} else {
			fmt.Printf("無此ID:%v帳號\n", id)
			continue
		}
	}
}

func (cv *customerView) exit() {
	choice := ""
	for {
		fmt.Printf("確定是否退出(Y/N):")
		fmt.Scanln(&choice)
		if choice == "Y" || choice == "y" {
			cv.loop = false
			break
		} else if choice == "N" || choice == "n" {
			break
		} else {
			fmt.Println("輸入錯誤，請重新輸入")
		}
	}
}

func (cv *customerView) mainMenu() {
	cv.loop = true
	for {
		cv.key = ""
		fmt.Println("--------------客戶資料管理系統--------------")
		fmt.Printf("\t\t1.增 加 客 戶\n")
		fmt.Printf("\t\t2.修 改 客 戶\n")
		fmt.Printf("\t\t3.刪 除 客 戶\n")
		fmt.Printf("\t\t4.客 戶 列 表\n")
		fmt.Printf("\t\t5.退 出 系 統\n")
		fmt.Printf("請選擇(1-5): ")
		fmt.Scanln(&cv.key)
		switch cv.key {
		case "1":
			cv.add()
		case "2":
			cv.edit()
		case "3":
			cv.delete()
		case "4":
			cv.list()
		case "5":
			cv.exit()
		default:
			fmt.Println("輸入錯誤，請重新輸入")
		}
		if cv.loop == false {
			break
		}
	}

	fmt.Println("退出客戶管理系統")
}

func main() {
	var customerView customerView
	customerView.customerService = service.NewCustomerService()
	customerView.mainMenu()
}
