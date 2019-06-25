package main

import (
	"fmt"
)

const (
	d = iota
	e = iota
	f = iota
)

func main() {
	// Bit shifting by 3
	a := 8              // 2^3
	fmt.Println(a << 3) // 2^3 * 2^3 =2^6=64
	fmt.Println(a >> 3) // 2^3 / 2^3 = 2^0=1
	s := "abcdefghijk"
	fmt.Printf("%v,%T\n", s, s)
	fmt.Printf("%v,%T\n", s[2], s[2]) // When slicing a string it gives unint8
	b := []byte(s)
	s1 := string(b)
	fmt.Printf("%v", b)      // Printing byte array
	fmt.Printf("\n%v\n", s1) // Printing string conversion of byte array
	fmt.Printf("%v\n", d)
	fmt.Printf("%v\n", e)
	fmt.Printf("%v\n", f)
}
