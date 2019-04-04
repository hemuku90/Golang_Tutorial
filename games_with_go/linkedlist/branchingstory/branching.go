package main

import (
	"bufio"
	"fmt"
	"os"
)

// Implementation of a Binary tree
type storyNode struct {
	text    string
	yesPath *storyNode
	noPath  *storyNode
}

func (node *storyNode) play() {
	fmt.Println(node.text)
	scanner := bufio.NewScanner(os.Stdin)
	if node.yesPath != nil && node.noPath != nil {
		for {
			scanner.Scan()
			answer := scanner.Text()
			if answer == "yes" {
				node.yesPath.play()
				break
			} else if answer == "no" {
				node.noPath.play()
				break
			} else {
				fmt.Println("Please select Yes or No")
			}
		}
	}
}
func main() {
	root := storyNode{"Your are at the entrance of the cave. Do you want to go in the cave?", nil, nil}
	winning := storyNode{"Winning", nil, nil}
	losing := storyNode{"Losing", nil, nil}
	root.yesPath = &losing
	root.noPath = &winning
	root.play()
}
