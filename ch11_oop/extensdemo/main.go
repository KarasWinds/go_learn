package main

import "fmt"

type student struct {
	name  string
	age   int
	score int
}

func (s *student) showInfo() {
	fmt.Printf("姓名:%v,年齡:%v,分數:%v\n", s.name, s.age, s.score)
}

func (s *student) setScore(score int) {
	s.score = score
}

type pupil struct {
	student
}

func (p *pupil) testing() {
	fmt.Println("小學生正在考試...")
}

type graduate struct {
	student
}

func (g *graduate) testing() {
	fmt.Println("大學生正在考試...")
}

func main() {
	pp := &pupil{}
	pp.name = "aki"
	pp.age = 10

	pp.testing()
	pp.setScore(80)
	pp.showInfo()

	gg := &graduate{}
	gg.name = "bin"
	gg.age = 20

	gg.testing()
	gg.setScore(90)
	gg.showInfo()

}
