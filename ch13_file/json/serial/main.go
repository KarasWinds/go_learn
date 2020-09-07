package main

import (
	"encoding/json"
	"fmt"
)

// Monster is a class
type Monster struct {
	Name     string  `json:"name"`
	Age      int     `json:"age"`
	Birthday string  `json:"birthday"`
	Sal      float64 `json:"sal"`
	Skill    string  `json:"skill"`
}

func testStruct() {
	monster := Monster{
		Name:     "九尾",
		Age:      999,
		Birthday: "0000-00-00",
		Sal:      9999.99,
		Skill:    "咆嘯",
	}
	fmt.Println(monster)
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("序列化錯誤, err:%v\n", err)
	}
	fmt.Printf("monster序列化後:%v\n", string(data))
}

func testMap() {
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "酒吞"
	a["age"] = 800
	a["skill"] = "嘔吐"
	fmt.Println(a)
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("序列化錯誤, err:%v\n", err)
	}
	fmt.Printf("monster序列化後:%v\n", string(data))
}

func testSlice() {
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "tom"
	m1["age"] = 22
	m1["skill"] = "speak"
	slice = append(slice, m1)

	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"] = "257"
	m2["age"] = 28
	m2["skill"] = "lulu"
	slice = append(slice, m2)
	fmt.Println(slice)
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("序列化錯誤, err:%v\n", err)
	}
	fmt.Printf("monster序列化後:%v\n", string(data))
}

func testFloat64() {
	var num float64 = 32.3
	fmt.Println(num)
	data, err := json.Marshal(&num)
	if err != nil {
		fmt.Printf("序列化錯誤, err:%v\n", err)
	}
	fmt.Printf("monster序列化後:%v\n", string(data))

}

func main() {
	testStruct()
	testMap()
	testSlice()
	testFloat64()
}
