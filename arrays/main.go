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
	// Slices
	slice := make([]int, 3)
	fmt.Printf("Slice %v", slice)
	fmt.Printf("Length is %v and capcity is %v\n", len(slice), cap(slice))
	num := []int{1, 2, 3, 4, 5}
	fmt.Println(num)
	num2 := append(num[:2], num[3:]...) // Removing middle element from a slice
	fmt.Println(num2)
	fmt.Println(num)
	// Pointer  to slice
	type num3 []int
	num3={1,2}
	var ptr *num3
	fmt.Println("Pointer\n", ptr)
}
