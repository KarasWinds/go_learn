package main

import "fmt"

func main() {
	fmt.Println(10 / 4)
	var n1 float32 = 10 / 4
	fmt.Println(n1)

	var n2 float32 = 10.0 / 4
	fmt.Println(n2)

	fmt.Println("10 % 3 =", 10%3)
	fmt.Println("-10 % 3 =", -10%3)
	fmt.Println("10 % -3 =", 10%-3)
	fmt.Println("-10 % -3 =", -10%-3)

	var i int = 10
	i++
	fmt.Println("i =", i)
	i--
	fmt.Println("i =", i)

	// X天等於幾周幾天
	var days int = 97
	var week int = 97 / 7
	var day int = 97 % 7
	fmt.Printf("%d天等於%d周%d天\n", days, week, day)

	// 定義一個變數保存華氏溫度，並轉換為攝氏溫度
	// 轉換公式為攝氏溫度=5/9*(華氏溫度-100)
	var huashi float32 = 154.5
	var sheshi float32 = 5.0 / 9 * (huashi - 100)
	fmt.Printf("華氏溫度%v等於攝氏溫度%v\n", huashi, sheshi)
}
