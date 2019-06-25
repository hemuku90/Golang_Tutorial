// Selection Sort
package main

import "fmt"

func insertArray() *[]int {
	var n int
	fmt.Println("Enter size of the array")
	fmt.Scan(&n)
	arr := make([]int, n)
	fmt.Println("Enter an array to be sorter")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	return &arr
}
func selectionSort(arr **[]int) []int {
	a := **arr
	n := len(a)
	fmt.Println(n)
	fmt.Println("From func:", a)
	for i := 0; i < n-1; i++ {
		var imin = i
		for j := i + 1; j < n; j++ {
			if a[j] < a[imin] {
				imin = j
			}
		}
		temp := a[i]
		a[i] = a[imin]
		a[imin] = temp
	}
	return a
}
func main() {
	arr := insertArray()
	sortedArray := selectionSort(&arr)
	fmt.Println(sortedArray)
}
