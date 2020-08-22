package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num1 int = 99
	var num2 float64 = 23.44
	var b bool = true
	var char byte = 'g'
	var str string
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type %T str=%q\n", str, str)
	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type %T str=%q\n", str, str)
	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type %T str=%q\n", str, str)
	str = fmt.Sprintf("%c", char)
	fmt.Printf("str type %T str=%q\n", str, str)

	var num3 int = 99
	var num4 float64 = 23.432
	var b2 bool = true
	str = strconv.FormatInt(int64(num3), 10)
	fmt.Printf("str type %T str=%q\n", str, str)
	str = strconv.FormatFloat(num4, 'f', 10, 64)
	fmt.Printf("str type %T str=%q\n", str, str)
	str = strconv.FormatBool(b2)
	fmt.Printf("str type %T str=%q\n", str, str)

	var num5 int64 = 203
	str = strconv.Itoa(int(num5))
	fmt.Printf("str type %T str=%q\n", str, str)
}
