package concepts

import (
	"fmt"
)

func Chan() {
	fmt.Println("Intro to Channels")
	c := make(chan int, 2) // defining a buffered channel
	//Writing to a channel
	c <- 10
	c <- 25
	fmt.Println("\nLength of a channel", len(c))
	// reading from a channel
	var v = <-c
	fmt.Println("v=%v\n", v)
	fmt.Println("\nLength of a channel", len(c))
}
