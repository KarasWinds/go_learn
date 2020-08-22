package main

import "fmt"

func main() {
	var heroes = [...]string{"蝙蝠俠", "超人", "綠巨人"}

	for index, value := range heroes {
		fmt.Printf("heroes[%v]:%v\n", index, value)
	}

	for _, value := range heroes {
		fmt.Println(value)
	}

}
