package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch1 := make(chan int) // Unbuffered channel
	for j := 1; j < 5; j++ {
		wg.Add(2)
		go func() {
			i := <-ch1 // Recieve data from channel
			fmt.Println(i)
			wg.Done()
		}()
		go func() {
			k := 42
			ch1 <- k // Sending data to a channel
			wg.Done()
		}()
	}
	wg.Wait()
	ch := make(chan int)
	wg.Add(2)
	// Read only channel
	go func(ch <-chan int) { // Reading from a channel
		for {
			// Checking for a condition if chnnel is closed or open
			if i, ok := <-ch; ok {
				fmt.Println("From Read only channel:->", i)
			} else {
				break
			}
		}
		wg.Done()
	}(ch)
	// Write only channel
	go func(ch chan<- int) { // Writing into channel
		ch <- 50 // Sending data to a channel
		ch <- 45
		close(ch) // Letting the channel know that we are done sending data and closing it
		wg.Done()
	}(ch)
	wg.Wait()
}
