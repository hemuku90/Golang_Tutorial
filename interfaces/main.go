package main

import (
	"fmt"
	"math"
)

func main() {
	d := Circle{radius: 5}
	measure(d)

}

type Geometry interface {
	area() float64
	perimeter() float64
}
type Circle struct {
	radius float64
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
