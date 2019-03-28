package main

import (
	"fmt"
	"strconv"
)

var (
	i int = 50
)

func main() {
	// Concept of shadowing
	fmt.Printf("From global variable: %v, %T \n", i, i)
	var i int = 40
	fmt.Printf("From local scope: %v, %T \n", i, i)
	var j float32
	j = float32(i) // Casting variable i as float32 and assigning it to variable j
	fmt.Printf("%v, %T \n", j, j)
	k := strconv.Itoa(i)
	fmt.Printf("%v, %T \n", k, k)
}
