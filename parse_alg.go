package main

import "fmt"
import "os"

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

//create method for Node class to be able to add children nodes
func (n *Node) add_child(childID, conditionID int) {
	n.children[conditionID] = childID
}

//define a Graph structure
type Graph struct {
	nodes map[int]Node
}

//function to create Action object
func NewAction() Action {
	return Action{parameters: map[string]string{}}
}

//function to create Condition object
func NewCondtion() Condition {
	return Condition{pattern: map[string]string{}}
}

// function to create Node object
func NewNode() Node {
	return Node{conditions: map[int]Condition{}, children: map[int]int{}}
}

//function to create graph object
func NewGraph() Graph {
	return Graph{nodes: map[int]Node{}}
}

func main() {
	//open json file
	jsonFile, _ := os.Open("Music--2018-08-10T09_20_31.449Z.json")
	//catch if no file exists
	if jsonFile == nil {
		fmt.Println("no file in current directory")
	}

	//close the file so that it can be parsed later
	defer jsonFile.Close()

	fmt.Println(jsonFile)

}
