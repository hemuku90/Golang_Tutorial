package main

import "fmt"

func main() {
	fmt.Println("Pointers")
	var a = 44
	var b = &a
	fmt.Println(a, *b, b, &a)
	a = 27
	fmt.Println(a, *b, b, &a)
	*b = 17 + *b
	fmt.Println(a, *b, b, &a)
	arr := []int32{1, 2, 3}
	arr0 := &arr[0]
	arr1 := &arr[1]
	arr2 := &arr[2]
	fmt.Printf("%p-%p-%p \n", arr0, arr1, arr2)
	fmt.Printf("Array values are:%v-%v-%v\n", *arr0, *arr1, *arr2)
}
