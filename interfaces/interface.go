package main

import "fmt"

type person struct {
	first  string
	last   string
	saying string
}

type secreAgent struct {
	person
	ltk bool // License to kill !!
}
type human interface {
	speak()
}

func (s secreAgent) speak() {
	fmt.Println(s.first, s.last, s.saying)
}
func (p person) speak() {
	fmt.Println(p.first, p.last, p.saying)
}
func foo(h human) {
	h.speak()
}

func main() {
	// Ploymorphism
	p1 := person{
		first:  "James",
		last:   "Bond",
		saying: "I am James Bond",
	}
	sa1 := secreAgent{
		person: person{
			first:  "Ian",
			last:   "Fleming",
			saying: "I am author of the book-James Bond !!!",
		},
		ltk: true,
	}
	foo(p1)
	foo(sa1)
}
