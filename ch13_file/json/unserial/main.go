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

func unmarshalStruct() {
	str := "{\"name\":\"九尾\",\"age\":999,\"birthday\":\"0000-00-00\",\"sal\":9999.99,\"skill\":\"咆嘯\"}"
	fmt.Printf("反序列化前:%v\n", str)
	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Printf("unmarshal err:%v", err)
	} else {
		fmt.Printf("反序列化後:%v\n", monster)
	}
}

func unmarshalMap() {
	str := "{\"age\":800,\"name\":\"酒吞\",\"skill\":\"嘔吐\"}"
	var a map[string]interface{}
	fmt.Printf("反序列化前:%v\n", str)
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Printf("unmarshal err:%v", err)
	} else {
		fmt.Printf("反序列化後:%v\n", a)
	}
}

func unmarshalSlice() {
	str := "[{\"age\":22,\"name\":\"tom\",\"skill\":\"speak\"},{\"age\":28,\"name\":\"257\",\"skill\":\"lulu\"}]"
	var slice []map[string]interface{}
	fmt.Printf("反序列化前:%v\n", str)
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Printf("unmarshal err:%v", err)
	} else {
		fmt.Printf("反序列化後:%v\n", slice)
	}
}

func main() {
	unmarshalStruct()
	unmarshalMap()
	unmarshalSlice()
}
