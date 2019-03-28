package main

import "fmt"

func main() {
	fmt.Println("variadic functions") // It is a function which takes n number of parameters
	fmt.Println(addition(5, 2, 5, 7))
}
func add(a, b int) int {
	return a + b
}

// Syntax to write variadic function
func addition(num ...int) int {
	var sum = 0
	for _, values := range num {
		sum += values
	}
	return sum
}
