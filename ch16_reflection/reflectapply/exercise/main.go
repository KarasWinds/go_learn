package main

import (
	"fmt"
	"reflect"
)

type call struct {
	Num1 int
	Num2 int
}

func (c call) GetSub(name string) {
	fmt.Printf("%v完成了減法, %d - %d = %d\n", name, c.Num1, c.Num2, c.Num1-c.Num2)
}

func main() {
	var (
		c       *call
		inValue []reflect.Value
	)
	elem := reflect.New(reflect.TypeOf(c).Elem()).Elem()
	elem.FieldByName("Num1").SetInt(8)
	elem.FieldByName("Num2").SetInt(3)
	inValue = make([]reflect.Value, 1)
	inValue[0] = reflect.ValueOf("tom")
	elem.Method(0).Call(inValue)

	num := elem.NumField()
	fmt.Printf("struct has %d fileds\n", num)
	numOfMethod := elem.NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethod)

}
