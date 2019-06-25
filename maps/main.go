package main

import "fmt"

func test() {
	//fmt.Println("Map")
	var nameMap = make(map[int]string) // to create a map use make
	nameMap[1] = "Hemant"
	nameMap[2] = "Mukul"
	delete(nameMap, 2) // delete(nameofmap, key)
	fmt.Println(nameMap)

}
