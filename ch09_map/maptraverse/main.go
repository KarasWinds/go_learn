package main

import "fmt"

func main() {
	cities := make(map[string]string)
	cities["no1"] = "台北"
	cities["no2"] = "台中"
	cities["no3"] = "台南"

	for k, v := range cities {
		fmt.Printf("key:%v,value:%v\n", k, v)
	}

	fmt.Printf("cities有%v個\n", len(cities))

	studentMap := make(map[string]map[string]string)

	studentMap["no1"] = make(map[string]string)
	studentMap["no1"]["name"] = "tom"
	studentMap["no1"]["sex"] = "男"

	studentMap["no2"] = make(map[string]string)
	studentMap["no2"]["name"] = "may"
	studentMap["no2"]["sex"] = "女"

	for k1, v1 := range studentMap {
		fmt.Printf("ID:%v\n", k1)
		for k2, v2 := range v1 {
			fmt.Printf("\t%v:%v\n", k2, v2)
		}
	}
}
