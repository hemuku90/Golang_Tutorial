package main

import "fmt"

func main() {
	/*
			fmt.Println("Loops")
			for x := 1; x <= 10; x++ {
				fmt.Println(x)
			}

		var age = 25
		if age <= 25 {
			fmt.Println("You are less than 25")
		} else {
			fmt.Println("You are either greater or smaller than 25")
		}
	*/
	var age = 31
	switch age {
	case 30:
		fmt.Println("Your age is 30")
	case 35:
		fmt.Println("Your age is 35")
	default:
		fmt.Println("none of the ages are present")
	}
}
