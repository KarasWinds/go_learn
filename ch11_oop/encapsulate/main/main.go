package main

import (
	"fmt"
	"go_learn/ch11_oop/encapsulate/model"
)

func main() {
	p := model.NewPerson("kai")
	p.SetAge(22)
	p.SetSal(40000.0)

	fmt.Println(p)
	fmt.Printf("名:%v,年齡:%v,薪水:%v\n", p.Name, p.GetAge(), p.GetSal())
}
