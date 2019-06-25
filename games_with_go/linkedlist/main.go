package main

import "fmt"

type List struct {
	head *Node
	tail *Node
}
type Node struct {
	// Node will hold a value and reference to the next node
	value int   // Node will hold a value
	next  *Node // Pointer to next node
}

func (l *List) First() *Node {
	return l.head
}
func (l *List) Push(value int) {
	node := &Node{value: value}
	if l.head == nil {
		l.head = node
	} else {
		l.tail.next = node
	}
	l.tail = node
}
func (n *Node) Next() *Node {
	return n.next
}

func main() {
	l := &List{}
	fmt.Println(*l)
	l.Push(1)
	l.Push(2)
	l.Push(3)
	n := l.First()
	fmt.Println(n.value)
	n = n.Next()
	println(n.value)
	n = n.Next()
	println(n.value)

}
