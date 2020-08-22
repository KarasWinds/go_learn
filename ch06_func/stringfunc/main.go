package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// golang 的編碼為UTF-8(ASCII的字元(字母、數字)佔一位元組，漢字佔三個位元組)
	str := "hello啦"
	fmt.Println("str len :", len(str))

	// 字元取出含中文
	str2 := "hello啦啦"
	r := []rune(str2)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字:%c\n", r[i])
	}

	// 字串轉int
	n, err := strconv.Atoi("234")
	if err != nil {
		fmt.Println("轉換錯誤 :", err)
	} else {
		fmt.Println("轉換成果 :", n)
	}

	// int轉字串
	str3 := strconv.Itoa(2345343)
	fmt.Printf("str=%v,str=%T\n", str3, str3)

	// 字元轉ASCII
	var bytes = []byte("hello")
	fmt.Printf("bytes=%v\n", bytes)

	// ASCII轉字元
	str4 := string([]byte{97, 98, 99})
	fmt.Printf("str=%v\n", str4)

	// 10進制轉2...進制
	str = strconv.FormatInt(123, 2)
	fmt.Printf("123的2進制:%v\n", str)
	str = strconv.FormatInt(123, 16)
	fmt.Printf("123的16進制:%v\n", str)

	// 尋找字串是否出現在指定字串中
	b := strings.Contains("Seafood", "foo")
	fmt.Printf("b=%v\n", b)

	// 計算一個字串中出現幾次指定字串
	num := strings.Count("ceheese", "e")
	fmt.Printf("num=%v\n", num)

	// 不區分大小寫字串比對
	b = strings.EqualFold("abc", "Abc")
	fmt.Printf("b=%v\n", b)
	fmt.Println("結果:", "abc" == "Abc")

	// 字串出現位置查找
	index := strings.Index("NTL_abc", "abc")
	fmt.Printf("index=%v\n", index)
	index = strings.LastIndex("go golang", "go")
	fmt.Printf("index=%v\n", index)

	// 字串替換(,數量-1表示全部)
	str = strings.Replace("go go hello", "go", "Taipei", 1)
	fmt.Printf("str=%v\n", str)

	// 字串轉陣列取出
	strArr := strings.Split("hello,world,ok", ",")
	fmt.Printf("strarr=%v\n", strArr)
	for i := 0; i < len(strArr); i++ {
		fmt.Printf("strArr[%v]=%v\n", i, strArr[i])
	}

	// 大小寫轉換
	str = "goLang Hello"
	str = strings.ToLower(str)
	fmt.Printf("str=%v\n", str)
	str = strings.ToUpper(str)
	fmt.Printf("str=%v\n", str)

	// 字串替換
	str = strings.TrimSpace(" hello,world   ")
	fmt.Printf("str=%q\n", str)
	str = strings.Trim("! hello!! ! ", "! ")
	fmt.Printf("str=%q\n", str)
	str = strings.TrimLeft("! hello!! ! ", "! ")
	fmt.Printf("str=%q\n", str)
	str = strings.TrimRight("! hello!! ! ", "! ")
	fmt.Printf("str=%q\n", str)

	// 檢查字串首尾
	b = strings.HasPrefix("http://10.1.1.1", "http://")
	fmt.Printf("b=%v\n", b)
	b = strings.HasSuffix("http://10.1.1.1", "http://")
	fmt.Printf("b=%v\n", b)

}
