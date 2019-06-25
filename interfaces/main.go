package main

import (
	"fmt"
	"math"
)

func main() {
	d := Circle{radius: 5}
	fmt.Println("Circle object:")
	measure(d)
	r := Rectangle{length: 10, width: 10}
	fmt.Println(("Rectangle object:"))
	measure(r)
}

type Geometry interface {
	area() float64
	perimeter() float64
}
type Circle struct {
	radius float64
}

type Rectangle struct {
	length float64
	width  float64
}

func (r Rectangle) perimeter() float64 {
	return 2 * (r.length + r.width)
}

func (r Rectangle) area() float64 {
	return r.length * r.width
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) perimeter() float64 {
	return math.Pi * c.radius * 2
}
func measure(g Geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}
