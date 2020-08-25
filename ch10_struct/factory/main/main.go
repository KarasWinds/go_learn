package main

import (
	"fmt"
	"go_learn/ch10_struct/factory/model"
)

func main() {
	var stu = model.NewStudent("tom", 90.3)

	fmt.Println(*stu)
	fmt.Println(stu.GetScore())
}
