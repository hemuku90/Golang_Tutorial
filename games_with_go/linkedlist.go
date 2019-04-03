package main

import (
	"fmt"
)

type person struct {
	age       int
	weight    int
	firstName string
	next      *person
}

func main() {
	mike := &person{33, 200, "mike", nil}
	personList := mike
	marshall := &person{56, 123, "marshall", nil}
	george := &person{60, 321, "George", nil}
	personList = addNodeEnd(marshall, personList)
	personList = addNodeEnd(george, personList)
	printList(personList)
	p := reverseRecurrsive(personList)
	fmt.Println(p)
}

// Printing the linked list
func printList(personList *person) {
	p := personList
	fmt.Println("Personlist", p)
	for ; p != nil; p = p.next {
		fmt.Println(p)
	}
}

// Appending to a LinkedList
func addNodeEnd(newPerson, personList *person) *person {
	if personList == nil {
		return personList
	}
	for p := personList; p != nil; p = p.next {
		if p.next == nil {
			p.next = newPerson
			return personList
		}
	}
	return personList
}

// Reversing a linked list
func reverseRecurrsive(personList *person) *person {
	if personList == nil {
		return personList
	}
	p := personList

	if p.next == nil {
		return p
	} else {
		newHead := reverseRecurrsive(p.next)
		p.next.next = p
		p.next = nil
		return newHead
	}
}
