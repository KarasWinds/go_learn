package main

import "fmt"

func getSum(n1 int, n2 int) int {
	return n1 + n2
}

func myFunc1(funvar func(int, int) int, num1 int, num2 int) int {
	return funvar(num1, num2)
}

type myFun func(int, int) int

func myFunc2(funvar myFun, num1 int, num2 int) int {
	return funvar(num1, num2)
}

func getSumAndSub(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

func sumtest(n1 int, args ...int) int {
	sum := n1
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

func main() {
	a := getSum
	fmt.Printf("a的數據類型%T,getSum的數據類型%T\n", a, getSum)
	res := a(1, 2)
	fmt.Println("res =", res)
	res2 := myFunc1(a, 20, 30)
	fmt.Println("res2 =", res2)

	type myInt int

	var num1 myInt
	var num2 int
	num1 = 40
	num2 = int(num1)
	fmt.Println("num1 =", num1, "num2 =", num2)

	res3 := myFunc2(a, 202, 304)
	fmt.Println("res3 =", res3)

	sum, sub := getSumAndSub(2, 5)
	fmt.Printf("sum=%v,sub=%v\n", sum, sub)

	res4 := sumtest(10, 23, 43, 312)
	fmt.Println("res4 =", res4)
}
