package model

import "fmt"

// Customer 顧客類別
type Customer struct {
	id     int
	name   string
	gender string
	age    int
	phone  string
	email  string
}

func NewCustomer(name string, gender string, age int, phone string, email string) Customer {
	return Customer{
		name:   name,
		gender: gender,
		age:    age,
		phone:  phone,
		email:  email,
	}
}

func (customer *Customer) SetID(id int) {
	customer.id = id
}

func (customer *Customer) SetName(name string) {
	customer.name = name
}

func (customer *Customer) SetGender(gender string) {
	customer.gender = gender
}

func (customer *Customer) SetAge(age int) {
	customer.age = age
}

func (customer *Customer) SetPhone(phone string) {
	customer.phone = phone
}
func (customer *Customer) SetEmail(email string) {
	customer.email = email
}

func (customer *Customer) GetId() int {
	return customer.id
}

func (customer *Customer) GetName() string {
	return customer.name
}

func (customer *Customer) GetGender() string {
	return customer.gender
}

func (customer *Customer) GetAge() int {
	return customer.age
}

func (customer *Customer) GetPhone() string {
	return customer.phone
}

func (customer *Customer) GetEmail() string {
	return customer.email
}

func (Customer *Customer) GetInfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v", Customer.id, Customer.name, Customer.gender, Customer.age, Customer.phone, Customer.email)
	return info
}
