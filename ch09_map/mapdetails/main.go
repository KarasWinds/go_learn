package main

import "fmt"

type student struct {
	Name string
	age  int
}

func modify(map1 map[int]int) {
	map1[10] = 9
}

func main() {
	map2 := make(map[int]int)
	map2[1] = 99
	map2[3] = 21
	map2[10] = 4
	fmt.Println(map2)
	modify(map2)
	fmt.Println(map2)

	students := make(map[string]student, 10)
	stu1 := student{"ann", 18}
	stu2 := student{"bal", 20}
	students["01"] = stu1
	students["02"] = stu2

	fmt.Println(students)

	for k, v := range students {
		fmt.Printf("學生編號:%v,姓名:%v,年齡:%v\n", k, v.Name, v.age)
	}

}
