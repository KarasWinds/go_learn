package service

import (
	"go_learn/ch12_project/customerManager/model"
)

// CustomerService 顧客服務類別
type CustomerService struct {
	customers   []model.Customer
	customerNum int
}

// NewCustomerService  建構子
func NewCustomerService() *CustomerService {
	var customerService CustomerService
	return &customerService
}

// CustomerList 取得顧客清單list
func (cs *CustomerService) CustomerList() []model.Customer {
	return cs.customers
}

// Add 新增顧客
func (cs *CustomerService) Add(customer model.Customer) bool {
	customer.SetID(cs.customerNum + 1)
	cs.customers = append(cs.customers, customer)
	cs.customerNum++
	return true
}

// Delete 刪除顧客
func (cs *CustomerService) Delete(id int) bool {
	var index int = cs.findByID(id)
	if index == -1 {
		return false
	}
	cs.customers = append(cs.customers[:index], cs.customers[index+1:]...)
	return true
}

// Check 確認顧客
func (cs *CustomerService) Check(id int) bool {
	var index int = cs.findByID(id)
	if index == -1 {
		return false
	}
	return true
}

// Get 取得顧客
func (cs *CustomerService) Get(id int) model.Customer {
	var index int = cs.findByID(id)
	return cs.customers[index]
}

// Update 更新顧客
func (cs *CustomerService) Update(id int, name string, gender string, age int, phone string, email string) bool {
	var index int = cs.findByID(id)
	if index == -1 {
		return false
	}
	if name != "" {
		cs.customers[index].SetName(name)
	}
	if gender != "" {
		cs.customers[index].SetGender(gender)
	}
	if age != 0 {
		cs.customers[index].SetAge(age)
	}
	if phone != "" {
		cs.customers[index].SetPhone(phone)
	}
	if email != "" {
		cs.customers[index].SetEmail(email)
	}
	return true
}

func (cs *CustomerService) findByID(id int) int {
	var index int = -1
	for i := 0; i < len(cs.customers); i++ {
		if cs.customers[i].GetId() == id {
			index = i
			break
		}
	}
	return index
}
