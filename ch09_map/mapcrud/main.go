package main

import "fmt"

func main() {

	cities := make(map[string]string)
	cities["no1"] = "台北"
	cities["no2"] = "台中"
	cities["no3"] = "台南"
	cities["no3"] = "高雄"
	fmt.Println(cities)

	delete(cities, "no3")
	delete(cities, "no4")
	fmt.Println(cities)

	val, ok := cities["no1"]
	if ok {
		fmt.Println(val)
	} else {
		fmt.Println("找不到")
	}

	cities = make(map[string]string)
	fmt.Println(cities)

}
