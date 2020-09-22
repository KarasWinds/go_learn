package main

import (
	"fmt"
	"reflect"
)

func reflect01(b interface{}) {
	rVal := reflect.ValueOf(b)

	fmt.Printf("rVal kind:%v\n", rVal.Kind())
	rVal.Elem().SetInt(20)
}

func main() {
	var num int = 10
	reflect01(&num)
	fmt.Println("num:", num)
}
