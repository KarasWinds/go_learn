package main

import (
	"encoding/json"
	"fmt"
)

// AA is XXX
type AA struct {
	Num int
}

// B is XXX
type B struct {
	Num int
}

// Monster is
type Monster struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var a AA
	var b B
	a = AA(b)
	fmt.Println(a, b)

	monster := Monster{"九尾", 999}
	jsonStr, err := json.Marshal(monster)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonStr))
}
