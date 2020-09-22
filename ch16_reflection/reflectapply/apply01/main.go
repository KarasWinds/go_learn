package main

import (
	"fmt"
	"reflect"
)

// Monster is a class
type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"monster_age"`
	Score float32
	Sex   string
}

func (s Monster) print() {
	fmt.Println("---start---")
	fmt.Println(s)
	fmt.Println("---end---")
}

func (s Monster) getSum(n1, n2 int) int {
	return n1 + n2
}

func (s Monster) set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

func testStruct(a interface{}) {
	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	kd := val.Kind()
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}
	num := val.NumField()
	fmt.Printf("struct has %d fields\n", num)
	for i := 0; i < num; i++ {
		fmt.Printf("Field [%d] value:%v\n", i, val.Field(i))
		tagVal := typ.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("Field [%d] tag:%v\n", i, tagVal)
		}
	}

	numOfMethod := val.NumMethod()
	// method list is sorted by method name (ASCII)
	fmt.Printf("struct has %d methods\n", numOfMethod)

	val.Method(1).Call(nil)

	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	// input []reflect.Vale, return []reflect.value
	res := val.Method(0).Call(params)
	fmt.Println("res=", res[0].Int())
}

func main() {
	var a Monster = Monster{
		Name:  "九尾",
		Age:   999,
		Score: 99999,
	}

	testStruct(a)
}
