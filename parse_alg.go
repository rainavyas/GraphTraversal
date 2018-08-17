package main

import "fmt"

//define an Action structure
type Action struct {
	typ        string
	command    string
	parameters map[string]string
}

//define a Condition structure
type Condition struct {
	id      int
	typ     string
	pattern map[string]string
}

// define a node structure
type Node struct {
	id         int
	typ        string
	label      string
	conditions map[int]Condition
	actions    []Action
	children   map[int]int
}

//function to create Action object
func NewAction() Action {
	return Action{parameters: map[string]string{}}
}

//function to create Condition object - necessary to not have nil pointer in initialised map
func NewCondtion() Condition {
	// return object of type condition and the map now initialised to be empty (not nil)
	return Condition{pattern: map[string]string{}}
}

// function to create Node object
func NewNode() Node {
	return Node{conditions: map[int]Condition{}}
}

func main() {

	fmt.Print("test")

	cond := NewCondtion()
	node := NewNode()
	action := NewAction()

	node.actions = append(node.actions, action)

	cond.id = 3
	node.id = 1

	node.conditions[2] = cond

	fmt.Print(node.actions[0])

}
