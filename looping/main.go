package main

import "fmt"

func main() {
	// Loops
	/*for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}*/
	for k := 0; k < 5; k++ {
		for l := 0; l < 5; l++ {
			fmt.Printf("%v\t", l*k)
		}
		fmt.Println("\n")
	}
}
