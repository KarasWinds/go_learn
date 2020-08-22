package main

import "fmt"

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return c.radius * c.radius * 3.14
}

func (c *circle) area2() float64 {
	fmt.Printf("area2() c位址:%p\n", c)
	return c.radius * c.radius * 3.14
}

func main() {
	var c1 circle
	c1.radius = 4.0
	res := c1.area()
	fmt.Println(res)

	var c2 circle
	fmt.Printf("main() c位址:%p\n", &c2)
	c2.radius = 5.0
	res = c2.area2()
	fmt.Println(res)
}
