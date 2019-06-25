package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{} // Waitgroup instance
var mu = sync.RWMutex{}
var counter = 0

func main() {
	runtime.GOMAXPROCS(10)
	for i := 0; i < 10; i++ {
		wg.Add(2)  // Telling the waitgroup to add 2 Goroutines
		mu.RLock() // Applying ReadLock on sayhello()
		go sayHello()
		mu.Lock() // Applying WriteLock on increment()
		go increment()
	}
	wg.Wait()
}

func sayHello() {
	fmt.Printf("Hello#%v\n", counter)
	mu.RUnlock()
	wg.Done()
}
func increment() {
	counter++
	mu.Unlock()
	wg.Done()
}
