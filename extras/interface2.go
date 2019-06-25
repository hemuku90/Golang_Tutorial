package main

import "fmt"

//Empty interace
type Shape interface{}

type Circle struct {
	radius float64
}

type Rectangle struct {
	length float64
	width  float64
}

func (r *Rectangle) Area(s interface{}) float64 {
	fmt.Printf("The value is %v and type is %T", s, s)
	return r.length * r.width
}
func (c *Circle) Area(s interface{}) float64 {
	fmt.Printf("The value is %v and type is %T", s, s)
	return c.radius * c.radius * 3.14
}

//main
func main() {
	var s Shape // Interface
	var circle = &Circle{radius: 14}
	s = circle
	fmt.Println("\n", s, circle.Area(s))
	var rectangle = &Rectangle{length: 10, width: 10}
	s = rectangle
	fmt.Println("\n", s, rectangle.Area(s))
}
