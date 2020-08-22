package main

import "fmt"

func getSum(n1 int, n2 int) int {
	fmt.Println(n1 + n2)
	return n1 + n2
}

func getSumAndSub(n1 int, n2 int) (int, int) {
	return n1 + n2, n1 - n2
}

func test(n1 int) {
	n1++
	fmt.Println(n1)
}

func main() {
	n1 := 1
	test(n1)
	fmt.Println(n1)
	sum := getSum(10, 20)
	fmt.Println(sum)

	res1, res2 := getSumAndSub(1, 3)
	fmt.Printf("res1=%v,res2=%v\n", res1, res2)

}
