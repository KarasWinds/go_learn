package main

import "fmt"

type student struct {
}

func typeJudge(items ...interface{}) {
	for i, v := range items {
		switch v.(type) {
		case bool:
			fmt.Printf("第%v個參數是%v，類型bool\n", i, v)
		case string:
			fmt.Printf("第%v個參數是%v，類型string\n", i, v)
		case int:
			fmt.Printf("第%v個參數是%v，類型int\n", i, v)
		case float64:
			fmt.Printf("第%v個參數是%v，類型float64\n", i, v)
		case student:
			fmt.Printf("第%v個參數是%v，類型student\n", i, v)
		case *student:
			fmt.Printf("第%v個參數是%v，類型*student\n", i, v)
		default:
			fmt.Printf("第%v個參數是%v，類型未知\n", i, v)
		}
	}
}

func main() {

	var n1 int = 1
	var n2 float32 = 2.2
	var n3 float64 = 2.4

	var stu student
	var stu2 *student

	typeJudge(n1, n2, n3, stu, stu2)

}
