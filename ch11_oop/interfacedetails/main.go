package main

import "fmt"

type aInterFace interface {
	say()
}

type bInterFace interface {
	hello()
}

type stu struct {
	name string
}

func (stu *stu) say() {
	fmt.Println("stu say()")
}

type integer int

func (i integer) say() {
	fmt.Println("integer say() ,i=", i)
}

type monster struct {
}

func (m monster) say() {
	fmt.Println("monster say()")
}

func (m monster) hello() {
	fmt.Println("monster hello()")
}

func main() {
	var stu1 stu
	var a aInterFace = &stu1
	a.say()

	var n integer = 22
	a = n
	a.say()

	var mon monster
	var a2 aInterFace = mon
	var b2 bInterFace = mon
	a2.say()
	b2.hello()

}
