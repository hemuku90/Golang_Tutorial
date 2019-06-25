package main

import "fmt"

type Queue struct {
	slice []int
}

// Push to a queue: At the end of a queue
func (q *Queue) enqueue(i int) {
	// i item to append to end of a queue
	q.slice = append(q.slice, i)
}

// Pop from a Queue: Returns first item from a Queue and removes that item from a Queue or panics if there isn't one

func (q *Queue) dequeue() int {
	// TO DO
	ret := q.slice[0] // First of our slice is the front of our queue
	q.slice = q.slice[1:len(q.slice)]
	return ret
}

func main() {
	var q *Queue = new(Queue)
	q.enqueue(123)
	q.enqueue(456)
	q.enqueue(678)
	fmt.Println("After push", *q)
	q.dequeue()
	fmt.Println("After queue pop:", *q)
}
