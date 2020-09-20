package main

import (
	"fmt"
	"reflect"
)

func reflectTest01(b interface{}) {

	rType := reflect.TypeOf(b)
	fmt.Printf("rType type:%T，value:%v\n", rType, rType)

	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal type:%T，value:%v\n", rVal, rVal)

	num := 2 + rVal.Int()
	fmt.Println(num)

	iV := rVal.Interface()
	num2 := iV.(int)
	fmt.Println(num2)
}

func reflectTest02(b interface{}) {

	rType := reflect.TypeOf(b)
	fmt.Printf("rType type:%T，value:%v\n", rType, rType)

	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal type:%T，value:%v\n", rVal, rVal)

	fmt.Println(rType.Kind(), rVal.Kind())

	iV := rVal.Interface()
	fmt.Printf("iV type:%T，value:%v\n", iV, iV)

	stu, ok := iV.(Student)
	if ok {
		fmt.Printf("student name:%v\n", stu.Name)
	}
}

// Student is a class
type Student struct {
	Name string
	Age  int
}

func main() {
	var num int = 100
	reflectTest01(num)

	stu := Student{
		Name: "tom",
		Age:  18,
	}
	reflectTest02(stu)
}
