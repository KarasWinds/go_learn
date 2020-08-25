package main

import "fmt"

// Student is a class exec
type Student struct {
	name   string
	gender string
	age    int
	id     int
	score  float64
}

func (student *Student) say() string {
	infoStr := fmt.Sprintf("student > name:%v,gender:%v,age:%v,id:%v,score:%v",
		student.name, student.gender, student.age, student.id, student.score)
	return infoStr
}

// Box is a class exec
type Box struct {
	lenth  float64
	width  float64
	height float64
}

func (box *Box) getVolumn() float64 {
	return box.lenth * box.width * box.height
}

// Visitor is a oop
type Visitor struct {
	Name string
	Age  int
}

func (visitor *Visitor) showPrice() {
	if visitor.Age >= 90 || visitor.Age < 5 {
		fmt.Printf("遊客姓名:%v,年齡:%v,不適合遊玩\n", visitor.Name, visitor.Age)
		return
	}
	if visitor.Age > 18 {
		fmt.Printf("遊客姓名:%v,年齡:%v,全票\n", visitor.Name, visitor.Age)
	} else {
		fmt.Printf("遊客姓名:%v,年齡:%v,半票\n", visitor.Name, visitor.Age)
	}
}

func main() {
	var stu = Student{
		name:   "ann",
		gender: "woman",
		age:    20,
		id:     2000,
		score:  80.88,
	}
	fmt.Println(stu.say())

	var box Box
	box.lenth = 1.2
	box.width = 2.2
	box.height = 4.3
	volumn := box.getVolumn()
	fmt.Printf("體積為:%v\n", volumn)

	var v Visitor
	for {
		fmt.Println("請輸入姓名:")
		fmt.Scanln(&v.Name)
		fmt.Println("請輸入年齡:")
		fmt.Scanln(&v.Age)
		if v.Name == "n" {
			break
		}
		v.showPrice()
	}

}
