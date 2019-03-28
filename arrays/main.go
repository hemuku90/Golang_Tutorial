package main

import "fmt"

func main() {
	fmt.Println("Arrays")
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
	//2-D array
	var identityMatrix = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	fmt.Println("\n", identityMatrix)
	var identityMatrix2 [3][3]int
	identityMatrix2[0] = [3]int{1, 0, 0}
	identityMatrix2[1] = [3]int{0, 1, 0}
	identityMatrix2[2] = [3]int{0, 0, 1}
	fmt.Println("\n", identityMatrix2)

}
