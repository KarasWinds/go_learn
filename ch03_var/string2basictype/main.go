package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string = "true"
	var b bool
	b, _ = strconv.ParseBool(str)
	fmt.Printf("b type %T b = %v\n", b, b)

	str = "12353"
	var n1 int64
	n1, _ = strconv.ParseInt(str, 10, 64)
	fmt.Printf("n1 type %T n1 = %v\n", n1, n1)

	str = "324.34"
	var f1 float64
	f1, _ = strconv.ParseFloat(str, 64)
	fmt.Printf("f1 type %T f1 = %v\n", f1, f1)
}
