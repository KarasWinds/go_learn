package main

import "fmt"

func main() {
	var monsters []map[string]string
	monsters = make([]map[string]string, 2)
	if monsters[0] == nil {
		monsters[0] = make(map[string]string)
		monsters[0]["name"] = "九尾"
		monsters[0]["age"] = "999"
	}
	if monsters[1] == nil {
		monsters[1] = make(map[string]string)
		monsters[1]["name"] = "酒吞童子"
		monsters[1]["age"] = "800"
	}
	fmt.Println(monsters)

	newMosters := map[string]string{
		"name": "忍野忍",
		"age":  "500",
	}

	monsters = append(monsters, newMosters)
	fmt.Println(monsters)
}
