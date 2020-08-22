package main

import "fmt"

type point struct {
	x int
	y int
}

type rect struct {
	leftUp, rightDown point
}

type rect2 struct {
	leftUp, rightDown *point
}

func main() {
	r1 := rect{point{1, 2}, point{3, 4}}

	fmt.Printf("r1.leftUp.x 位址:%p\nr1.leftUp.y 位址:%p\n", &r1.leftUp.x, &r1.leftUp.y)
	fmt.Printf("r1.rightDown.x 位址:%p\nr1.rightDown.y 位址:%p\n", &r1.rightDown.x, &r1.rightDown.y)

	r2 := rect2{&point{10, 20}, &point{30, 40}}

	fmt.Printf("r2.leftUp 本身位址:%p\nr2.rightDown 本身位址:%p\n", &r2.leftUp, &r2.rightDown)
	fmt.Printf("r2.leftUp 指向位址:%p\nr2.rightDown 指向位址:%p\n", r2.leftUp, r2.rightDown)
}
