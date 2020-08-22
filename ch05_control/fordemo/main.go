package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println("Hello~", i)
	}

	j := 1
	for j <= 10 {
		fmt.Println("World!", j)
		j++
	}

	k := 1
	for {
		if k <= 10 {
			fmt.Println("Ya", k)
		} else {
			break
		}
		k++
	}

	var str string = "hello,world!啦啦"
	str2 := []rune(str) //轉換為切片
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c \n", str2[i])
	}

	str = "for~test!酷酷"
	for index, val := range str {
		fmt.Printf("index=%d, val=%c\n", index, val)
	}
}
