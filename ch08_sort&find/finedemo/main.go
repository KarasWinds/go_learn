package main

import "fmt"

func main() {
	names := [4]string{"鷹", "獅", "龍", "虎"}
	var searchName string
	fmt.Println("請入要查找的名稱:")
	fmt.Scanln(&searchName)
	for i := 0; i < len(names); i++ {
		if searchName == names[i] {
			fmt.Printf("找到%v,在%v位置\n", names[i], i)
			break
		} else if i == len(names)-1 {
			fmt.Printf("沒有找到%v\n", searchName)
		}
	}

	index := -1
	for i := 0; i < len(names); i++ {
		if searchName == names[i] {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Printf("沒有找到%v\n", searchName)
	} else {
		fmt.Printf("找到%v,在%v位置\n", names[index], index)
	}

}
