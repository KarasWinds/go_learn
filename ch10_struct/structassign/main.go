package main

import "fmt"

// Stu is a class
type Stu struct {
	Name string
	Age  int
}

func main() {
	var stu1 = Stu{"tom", 12}
	stu2 := Stu{"ann", 20}

	var stu3 = Stu{
		Name: "kai",
		Age:  23,
	}

	stu4 := Stu{
		Name: "jkl",
		Age:  20,
	}

	fmt.Println(stu1, stu2, stu3, stu4)

	var stu5 = &Stu{"kni", 19}
	stu6 := &Stu{"qui", 22}

	var stu7 = &Stu{
		Name: "wei",
		Age:  28,
	}

	stu8 := &Stu{
		Name: "bei",
		Age:  17,
	}

	fmt.Println(*stu5, *stu6, *stu7, *stu8)
}
