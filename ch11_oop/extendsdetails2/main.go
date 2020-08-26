package main

import "fmt"

// AA is a class
type AA struct {
	Name string
	age  int
}

// BB is a class
type BB struct {
	Name string
	age  int
}

// CC is a class
type CC struct {
	niAA AA
	BB
	age int
}

type good struct {
	name  string
	price float64
}

type brand struct {
	name    string
	address string
}

type tv1 struct {
	good
	brand
}

type tv2 struct {
	*good
	*brand
}

type monster struct {
	name string
	age  int
}

type mm struct {
	monster
	int
	nn int
}

func main() {
	var exc CC
	exc.age = 18
	exc.niAA.Name = "tai"
	exc.BB.Name = "gu"
	fmt.Println(exc)

	TV1 := tv1{good{"tv", 1199.9}, brand{"hern", "台中"}}
	TV2 := tv1{
		good{
			name:  "tv",
			price: 1199.9,
		},
		brand{
			name:    "hern",
			address: "台中",
		},
	}
	fmt.Println(TV1)
	fmt.Println(TV2)

	TV3 := tv2{&good{"tv", 1199.9}, &brand{"hern", "台中"}}
	TV4 := tv2{
		&good{
			name:  "tv",
			price: 1199.9,
		},
		&brand{
			name:    "hern",
			address: "台中",
		},
	}
	fmt.Println(*TV3.good, *TV3.brand)
	fmt.Println(*TV4.good, *TV4.brand)

	var mon mm
	mon.name = "九尾"
	mon.age = 999
	mon.int = 1
	mon.nn = 9
	fmt.Println(mon)
}
