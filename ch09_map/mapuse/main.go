package main

import "fmt"

func main() {
	var name map[string]string
	name = make(map[string]string)
	name["no1"] = "tom"
	name["no2"] = "cat"
	fmt.Println(name)

	cities := make(map[string]string)
	cities["no1"] = "台北"
	cities["no2"] = "台中"
	fmt.Println(cities)

	heroes := map[string]string{
		"no1": "蝙蝠俠",
		"no2": "蜘蛛人",
	}
	fmt.Println(heroes)

	studentMap := make(map[string]map[string]string)

	studentMap["no1"] = make(map[string]string)
	studentMap["no1"]["name"] = "tom"
	studentMap["no1"]["sex"] = "男"

	studentMap["no2"] = make(map[string]string)
	studentMap["no2"]["name"] = "may"
	studentMap["no2"]["sex"] = "女"

	fmt.Println(studentMap)
	fmt.Println(studentMap["no2"])

}
