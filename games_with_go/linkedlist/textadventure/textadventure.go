package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Importing scanner
var scanner *bufio.Scanner

// Linked List
type choices struct {
	cmd         string
	description string
	nextNode    *storyNode
	nextChoice  *choices
}

// Graph with a Linked List choices
type storyNode struct {
	text    string
	choices *choices
}

// Function to add choices to our storyNode
func (node *storyNode) addChoice(cmd string, description string, nextNode *storyNode) {
	choice := &choices{cmd, description, nextNode, nil}
	if node.choices == nil {
		node.choices = choice
	} else {
		currentChoice := node.choices
		for currentChoice.nextChoice != nil {
			currentChoice = currentChoice.nextChoice
		}
		currentChoice.nextChoice = choice
	}
}

// Printing state of the node along with the choices
func (node *storyNode) render() {
	fmt.Println(node.text)
	currentChoice := node.choices
	for currentChoice != nil {
		fmt.Println(currentChoice.cmd, ":", currentChoice.description)
		currentChoice = currentChoice.nextChoice
	}
}

// Function to enter a command and send him to next storyNode
func (node *storyNode) executeCmd(cmd string) *storyNode {
	currentChoice := node.choices
	for currentChoice != nil {
		if strings.ToLower(currentChoice.cmd) == strings.ToLower(cmd) {
			return currentChoice.nextNode
		}
		currentChoice = currentChoice.nextChoice
	}
	fmt.Println("Choice did not match")
	return node
}

// Func to play on storyNode

func (node *storyNode) play() {
	node.render()
	if node.choices != nil {

		scanner.Scan()
		node.executeCmd(scanner.Text()).play()

	}
}
func main() {
	scanner = bufio.NewScanner(os.Stdin)
	start := storyNode{text: `
	You are in a large chamber deep underground.
	You see three passages. North passage leads into darkness.
	To the south it appears to lead out.`}
	darkRoom := storyNode{text: `It is pitch black. You cannot see a thing`}
	darkRoomLit := storyNode{text: `The dark passage is now lit by your lantern. You can continue head back north.`}
	grue := storyNode{text: "While stumbling you are eaten by a grue."}
	trap := storyNode{text: "You head down and suddenly trap door appears"}
	treassure := storyNode{text: "You find a treassure"}
	start.addChoice("N", "Go North", &darkRoom)
	start.addChoice("S", "Go South", &darkRoom)
	start.addChoice("E", "Go East", &trap)
	darkRoom.addChoice("S", "Try to go back South", &grue)
	darkRoom.addChoice("O", "Turn on your lantern", &darkRoomLit)
	darkRoomLit.addChoice("N", "Go get the treassure", &treassure)
	darkRoomLit.addChoice("S", "Go back to the start", &start)
	start.play()
	fmt.Println()
	fmt.Println("End of Game")
}
