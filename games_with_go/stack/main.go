package main

import "fmt"

type stack struct {
	slice []int
}

// Push at top of the stack
func (s *stack) push(i int) {
	s.slice = append(s.slice, i) // Add items at the top/end of the stack/queue respectively
}

// Returns top item from a stack but does not remove it
func (s *stack) peek() int {
	// peek an item from the stack
	return s.slice[len(s.slice)-1]
}

// Pop items from top of the stack and return it
func (s *stack) pop() int {
	var ret int = s.peek()
	s.slice = s.slice[0 : len(s.slice)-1]
	return ret
}
func main() {
	var s *stack = new(stack)
	s.push(1)
	s.push(2)
	s.push(3)
	fmt.Println(*s)
	s.pop()
	s.pop()
	s.pop()
	s.pop()
	s.pop()
	s.pop()
	fmt.Println(*s)
}
