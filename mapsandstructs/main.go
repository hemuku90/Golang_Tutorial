package main

import (
	"fmt"
)

type Doctor struct {
	number     int
	name       string
	companions []string
}

type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal // Embedding Animal struct into Bird
	Speed  float32
	CanFly bool
}

func main() {
	// Manipulation on Maps are done by reference on original underlying Map
	fmt.Println("Hello Maps and Structs")
	statePopulations := map[string]int{
		"California": 92324,
		"Texas":      12342,
		"Florida":    12387,
	}
	statePopulations["Georgia"] = 324324
	fmt.Println(statePopulations)
	fmt.Println(statePopulations["California"] == statePopulations["Florida"])
	delete(statePopulations, "Georgia")
	fmt.Println(statePopulations)
	_, ok := statePopulations["Florid"] // Convention of OK nothing concrete as to check if key is present in a map
	fmt.Println(ok)
	// Structure object of type Doctor
	doctor1 := Doctor{
		number: 1,
		name:   "Hemant",
		companions: []string{
			"Test1",
			"Test2",
		},
	}
	fmt.Println(doctor1)
	b := Bird{}
	b.Name = "Emu"
	b.CanFly = false
	b.Origin = "Australia"
	b.Speed = 148
	fmt.Println("Structure Bird output is \n", b)
}
