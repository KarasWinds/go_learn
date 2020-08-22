package main

import "fmt"

// Cat is ... (struct 建議於上方增加註解)
type Cat struct {
	Name  string
	Age   int
	Color string
}

func main() {
	var cat1 Cat
	fmt.Printf("%v\n", cat1)
	cat1.Name = "tom"
	cat1.Age = 3
	cat1.Color = "白"
	fmt.Printf("%v\n", cat1)
	fmt.Println(cat1.Name, cat1.Age, cat1.Color)
}
