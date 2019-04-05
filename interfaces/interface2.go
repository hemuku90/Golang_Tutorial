package main

import "fmt"

type Guitarist interface {
	// PlayGuitar prints out "Playing Guitar"
	// to the terminal
	PlayGuitar()
}

type BaseGuitarist struct {
	Name string
}

type AcousticGuitarist struct {
	Name string
}

func (b BaseGuitarist) PlayGuitar() {
	fmt.Printf("%s plays the Bass Guitar\n", b.Name)
}

func (b AcousticGuitarist) PlayGuitar() {
	fmt.Printf("%s plays the Acoustic Guitar\n", b.Name)
}

func main() {
	var player1 BaseGuitarist
	player1.Name = "Paul"
	player1.PlayGuitar()
	var player2 AcousticGuitarist
	player2.Name = "Ringo"
	var guitarists []Guitarist
	guitarists = append(guitarists, player1)
	guitarists = append(guitarists, player2)
	fmt.Println(guitarists)
}
