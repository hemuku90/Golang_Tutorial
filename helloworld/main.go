package main

import "fmt"

func main() {
	var name string = "Hemant"
	fmt.Println(name)
	var lname = "Kumar"
	fmt.Println(lname)
	age := 28
	fmt.Println(age)
	// Multiple variables declaration
	var (
		lang   = "English"
		gender = "Male"
	)
	fmt.Println(lang, gender)
	const pi = 3.14
	fmt.Println(pi)
}
