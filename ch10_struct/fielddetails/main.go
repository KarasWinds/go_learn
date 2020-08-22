package main

import "fmt"

type person struct {
	Name   string
	Age    int
	Scores [5]float64
	ptr    *int
	slice  []int
	map1   map[string]string
}

type monster struct {
	Name string
	Age  int
}

func main() {
	var p1 person
	fmt.Println(p1)

	if p1.ptr == nil {
		fmt.Println("ok1")
	}
	if p1.slice == nil {
		fmt.Println("ok2")
	}
	if p1.map1 == nil {
		fmt.Println("ok3")
	}

	p1.slice = make([]int, 3)
	p1.slice[0] = 100
	fmt.Println(p1)

	p1.map1 = make(map[string]string)
	p1.map1["key"] = "value"
	fmt.Println(p1)

	var monster1 monster
	monster1.Name = "九尾"
	monster1.Age = 999

	monster2 := monster1
	monster2.Name = "酒吞"

	fmt.Println(monster1)
	fmt.Println(monster2)

}
