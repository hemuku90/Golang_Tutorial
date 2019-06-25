package main

import (
	"errors"
	"fmt"
)

var ErrInvalidNode = errors.New("Invalid Node")

// Node Interface
type Node interface {
	SetValue(v int) error
	GetValue() int
}

// Sll Node
type SllNode struct {
	next         *SllNode
	value        int
	SNodeMessage string
}

//Method SetValue for SllNode
func (sNode *SllNode) SetValue(v int) error {
	if sNode == nil {
		return ErrInvalidNode
	}
	sNode.value = v
	return nil
}

// Method GetValue for SllNode
func (sNode *SllNode) GetValue() int {
	return sNode.value
}

// Constructor for SllNode
func NewSllNode() *SllNode {
	// Initialize the struct and return its default value
	return &SllNode{SNodeMessage: "This is a message for the normal node"}
}

//type PowerNode
type PowerNode struct {
	next         *PowerNode
	value        int
	PNodeMessage string
}

//method SetValuefor Powernode
func (pNode *PowerNode) SetValue(v int) error {
	if pNode == nil {
		return ErrInvalidNode
	}
	pNode.value = v * 10
	return nil
}

//method Getvalue for Powernode
func (pNode *PowerNode) GetValue() int {
	return pNode.value
}

// constructor for Powernode
func NePowerNode() *PowerNode {
	return &PowerNode{PNodeMessage: "This is a message for the Power Node"}
}

// intialise the node
func createNode(v int) Node {
	sn := NewSllNode()
	sn.SetValue(v)
	return sn
}

//main
func main() {
	n := createNode(5)
	n1 := createNode(6)
	// switch concreten := n.(type) {
	// case *SllNode:
	// 	fmt.Println("Type is  SllNode, message: ", concreten.SNodeMessage)
	// case *PowerNode:
	// 	fmt.Println(" Type is  PowerNode, message: ", concreten.PNodeMessage)
	// }
	fmt.Println("This is stuct :", n, n1)
}
