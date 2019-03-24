package main

import "fmt"

func main() {
	fmt.Println("Arrays")
	// array
	var nums [2]int32
	nums[0] = 0
	nums[1] = 2
	fmt.Print(nums, "\n")
	newnums := [2]int{3, 4} // Second way to create an array
	fmt.Println(newnums)
	for pos, values := range newnums {
		fmt.Println(pos, values)
	}
	// Slices are dynamic array which does not contain size while declaring
	name := []string{"Hemant", "Pushpender", "Mukul"}
	for _, values := range name {
		fmt.Println(values)
	}
}
