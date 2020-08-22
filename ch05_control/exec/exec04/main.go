package main

import "fmt"

func main() {
	// var char byte
	// fmt.Println("請輸入一個字母:")
	// fmt.Scanf("%c", &char)

	// switch char {
	// case 'a':
	// 	fmt.Println("A")
	// case 'b':
	// 	fmt.Println("B")
	// case 'c':
	// 	fmt.Println("C")
	// case 'd':
	// 	fmt.Println("D")
	// case 'e':
	// 	fmt.Println("E")
	// default:
	// 	fmt.Println("other")
	// }

	// var score float32
	// fmt.Println("請輸入成績:")
	// fmt.Scanln(&score)

	// switch {
	// case score > 100:
	// 	fmt.Println("輸入有誤")
	// case score >= 60:
	// 	fmt.Println("及格")
	// case score >= 0:
	// 	fmt.Println("不及格")
	// default:
	// 	fmt.Println("輸入有誤")
	// }

	// var month byte
	// fmt.Println("請輸入月份:")
	// fmt.Scanln(&month)
	// switch month {
	// case 3, 4, 5:
	// 	fmt.Println("春天")
	// case 6, 7, 8:
	// 	fmt.Println("夏天")
	// case 9, 10, 11:
	// 	fmt.Println("秋天")
	// case 12, 1, 2:
	// 	fmt.Println("冬天")
	// default:
	// 	fmt.Println("輸入錯誤")
	// }

	var week string
	fmt.Println("請輸入星期:")
	fmt.Scanln(&week)

	switch week {
	case "星期一":
		fmt.Println("今天星期一")
	case "星期二":
		fmt.Println("今天星期二")
	case "星期三":
		fmt.Println("今天星期三")
	case "星期四":
		fmt.Println("今天星期四")
	case "星期五":
		fmt.Println("今天星期五")
	case "星期六":
		fmt.Println("今天星期六")
	case "星期日":
		fmt.Println("今天星期日")
	default:
		fmt.Println("輸入錯誤")
	}
}
