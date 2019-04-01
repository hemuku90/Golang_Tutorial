package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	for j := 1; j < 5; j++ {
		wg.Add(2)
		go func() {
			i := <-ch // Recieving data from channel
			fmt.Println(i)
			wg.Done()
		}()
		go func() {
			k := 42
			ch <- k // Sending data to a channel
			wg.Done()
		}()
	}
	wg.Wait()
}
