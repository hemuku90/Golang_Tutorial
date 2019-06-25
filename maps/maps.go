package main

import "fmt"

// Package level variable
var nameMap = make(map[string]string)

func init() {
	nameMap["Name1"] = "Hemant"
	nameMap["Name2"] = "Madhur"
}

// Update from a map
func update(key string, value string) {
	nameMap[key] = value
}

// Delete from a map
func deleteMap(key string) {
	// Check if map has the key
	if _, ok := nameMap[key]; ok {
		delete(nameMap, key)
	} else {
		fmt.Println("this key is not present", key)
	}
}

func main() {
	fmt.Println("#######map is##########\n", nameMap)
	// adding another key
	update("Name3", "Harjeet")
	fmt.Println("#######Updated a map#### \n", nameMap)
	// Deleting from a map
	deleteMap("Name2")
	fmt.Println("######New map after deletion#####\n", nameMap)
}
