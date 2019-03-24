package main

import "fmt"

func main() {
	fmt.Println("Hello")
	//fmt.Println(add(2, 3))
	//fmt.Println(sub(3, 2))
	add, sub := addsub(10, 3)
	fmt.Println(add, sub)
	//	fmt.Println(addsub(5, 3))
}
func add(a, b int) int {
	var sum = a + b
	return sum
}
func sub(a, b int) int {
	var sub = a - b
	return sub
}
func addsub(a, b int) (int, int) {
	return a + b, a - b
}
