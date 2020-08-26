package main

import "fmt"

type monkey struct {
	name string
}

type birdAble interface {
	flying()
}

type fishAble interface {
	swimming()
}

func (mon *monkey) climbing() {
	fmt.Println(mon.name, "爬樹")
}

type littleMonkey struct {
	monkey
}

func (mon *littleMonkey) flying() {
	fmt.Println(mon.name, "飛翔")
}

func (mon *littleMonkey) swimming() {
	fmt.Println(mon.name, "游泳")
}

func main() {
	var monkey littleMonkey
	monkey.name = "悟空"
	monkey.climbing()
	monkey.flying()
	monkey.swimming()
}
