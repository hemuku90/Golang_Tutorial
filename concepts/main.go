package main

import (
	"fmt"

	channels "Golang_Tutorial/concepts/channels"
)

func main() {
	var slice = []int{1, 2, 3, 4, 5}
	var ptr = &slice
	var ptr1 []*int
	//fmt.Print((*ptr)[0])
	for i := range *ptr {
		ptr1 = append(ptr1, &((*ptr)[i]))
	}
	fmt.Println("This is ptr:\n", ptr1) // Slice of pointers to slice
	for _, value := range ptr1 {
		fmt.Println(*value)
	}
	fmt.Println("About Channels:")
	channels.Chan()
}
