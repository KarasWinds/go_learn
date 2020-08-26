package main

import "fmt"

type bInterFace interface {
	test01()
}

type cInterFace interface {
	test02()
}

type aInterFace interface {
	bInterFace
	cInterFace
	test03()
}

type student struct {
}

func (stu student) test01() {
	fmt.Println("test01...")
}
func (stu student) test02() {
	fmt.Println("test02...")
}
func (stu student) test03() {
	fmt.Println("test03...")
}

type tt interface {
}

func main() {
	var stu student
	var a aInterFace = stu
	a.test01()
	a.test02()
	a.test03()

	var tt tt
	tt = stu
	fmt.Println(tt)

	var t2 interface{} = stu
	var num1 float64 = 8.8
	t2 = num1
	tt = num1
	fmt.Println(t2, tt)
}
