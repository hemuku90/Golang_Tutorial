package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	wg.Add(1)
	go sayHello()
	wg.Add(1)
	var msg = "World"
	go func() {
		fmt.Println(msg)
		wg.Done()
	}()
	wg.Wait()
}
func sayHello() {
	fmt.Println("Hello")
	wg.Done()
}
