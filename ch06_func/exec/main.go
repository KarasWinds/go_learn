package main

import (
	"fmt"
	"math/rand"
	"time"
)

func exec1() {
	var year int
	var month int8
	for {
		fmt.Println("請輸入年份")
		fmt.Scanln(&year)
		fmt.Println("請輸入月份")
		fmt.Scanln(&month)
		switch month {
		case 1, 3, 5, 7, 8, 10, 12:
			fmt.Println("天數31")
		case 4, 6, 9, 11:
			fmt.Println("天數30")
		case 2:
			if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
				fmt.Println("天數29")
			} else {
				fmt.Println("天數28")
			}
		default:
			fmt.Println("輸入錯誤")
			continue
		}
	}
}

func exec2() {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(100) + 1
	fmt.Println("輸入猜測數字(1~100):", num)
	var ans int
	for i := 1; i <= 10; i++ {
		fmt.Printf("第%v次輸入猜測數字(1~100):", i)
		fmt.Scanln(&ans)
		if ans == num {
			if i == 1 {
				fmt.Println("真是天才")
				break
			} else if i <= 3 {
				fmt.Println("不錯喔")
				break
			} else if i <= 9 {
				fmt.Println("普普通通")
				break
			}
			fmt.Println("總算猜對啦")
		}

	}
	if ans != num {
		fmt.Println("竟然都沒有猜對")
	}

}

func exec3() {
	var sum int
	var len int
	for i := 2; i <= 100; i++ {
		for j := 2; j <= i; j++ {
			if i == j {
				fmt.Printf("%v\t", i)
				len++
				sum += i
				if len%5 == 0 {
					fmt.Println("sum =", sum)
					sum = 0
				}
			} else if i%j == 0 {
				break
			}
		}
	}
}

// 從1990-01-01開始三天打魚兩天曬網
func exec4() {
	a, _ := time.Parse("2006-01-02", "1990-01-01")
	b, _ := time.Parse("2006-01-02", "2020-01-06")
	c := int(b.Sub(a).Hours()/24) + 1
	if c%5 <= 3 {
		fmt.Println("打魚")
	} else {
		fmt.Println("曬網")
	}
}

func exec5() {
	var n1 int = 10
	var n2 int = 5
	for {
		var input int = -1
		fmt.Println("輸入1.加法、2.減法、3.乘法、4.除法、0.退出")
		fmt.Scanln(&input)
		switch input {
		case 1:
			fmt.Printf("%v + %v = %v\n", n1, n2, n1+n2)
		case 2:
			fmt.Printf("%v - %v = %v\n", n1, n2, n1-n2)
		case 3:
			fmt.Printf("%v * %v = %v\n", n1, n2, n1*n2)
		case 4:
			fmt.Printf("%v / %v = %v\n", n1, n2, n1/n2)
		case 0:
			fmt.Println("退出")
			goto end
		default:
			fmt.Println("輸入錯誤")
		}
	}
end:
}

func exec6() {
	var upeng = 65
	var loeng = 97
	for i := 0; i < 26; i++ {
		fmt.Printf("%c", upeng+i)
	}
	fmt.Println()
	for i := 0; i < 26; i++ {
		fmt.Printf("%c", loeng+i)
	}
	fmt.Println()

}

func main() {
	// exec1()
	// exec2()
	// exec3()
	// exec4()
	// exec5()
	// exec6()
}
