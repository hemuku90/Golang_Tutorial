package main

import (
	"fmt"
)

func main() {
	// Break statement is imlicitly callled in switch statements in Go
	i := 1
	switch {
	case i <= 1:
		fmt.Println("Number less than or equal to 1")
		fallthrough // To check if first condition passes then second one obviously passes true
	case i <= 2:
		fmt.Println("Number greater than 2 and less than 10")
	default:
		fmt.Println("Enter another number")
	}
	// Type switchin
	var j interface{} = "Hello" // Interface type can take any type variable int,string,float etc
	switch j.(type) {
	case int:
		fmt.Println("j is an integer")
	case float32:
		fmt.Println("j is float32")
	case string:
		fmt.Println("j is a string")
	default:
		fmt.Println("Another type")
	}
}
