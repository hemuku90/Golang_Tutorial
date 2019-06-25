package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func (v *Vertex) test() {
	fmt.Println("Hello")
	fmt.Println(v.X, v.Y)
}

func (v Vertex) test2() {
	fmt.Println("Hello")
}
func main() {
	v := Vertex{1, 2}
	p := &v
	p.test()
	fmt.Println(v, p)
	f := v
	f.X = 3
	fmt.Println(f.X)

}
