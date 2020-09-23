package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type monster struct {
	Name  string `json:"monster_name"`
	Age   int
	Score float32
	Sex   string
}

func (s monster) Print() {
	fmt.Println("---start---")
	fmt.Println(s)
	fmt.Println("---end---")
}

func testStruct(a interface{}) {
	tye := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	kd := val.Kind()
	if kd != reflect.Ptr && val.Elem().Kind() == reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	num := val.Elem().NumField()
	val.Elem().Field(0).SetString("369")
	for i := 0; i < num; i++ {
		fmt.Printf("%d %v\n", i, val.Elem().Field(i).Kind())
	}

	fmt.Printf("struct has %d fileds\n", num)

	tag := tye.Elem().Field(0).Tag.Get("json")
	fmt.Printf("tag=%s\n", tag)

	numOfMethod := val.Elem().NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethod)

	val.Elem().Method(0).Call(nil)
}

func main() {
	var a monster = monster{
		Name:  "333",
		Age:   999,
		Score: 99.9,
	}

	result, _ := json.Marshal(a)
	fmt.Println("json result:", string(result))

	testStruct(&a)
	fmt.Println(a)
}
