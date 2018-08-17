package main

import (
	"encoding/json"
	"fmt"

	"io/ioutil"
)

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

//define a set of structures required for json unmarsheling

//define json overall tree
type json_nodes_and_trans struct {
	nodes       []json_node       `json:"nodes"`
	transitions []json_transition `json:"transitions"`
}

//define the json node structure
type json_node struct {
	id    int    `json:"id"`
	typ   string `json:"type"`
	label string `json:"label`
	position json_position `json:"position"`
	conditions []json_condition `json:"conditions"`
	actions []json_action `json:"actions"`

}

//define the json position structure
type json_position{
	x int `json:"x"`
	y int `json:"y"`
}

//define the json condition structure
type json_condition struct{
	id int `json:"id"`
	typ int `json:"type"`
	pattern []json_pattern `json:"pattern"`
}

//define the json pattern structure
type json_pattern struct{
	key string `json:"key"`
	value string `json:"value"`

}

//define the json action structure
type json_action struct{
	typ string `json:"type"`
	command string `json:"command"`
	parameters []json_parameter `json:"parameters"`
}


//define the json paramater structure
type json_parameter struct{
	key string `json:"key"`
	value string `json"value"`
}

//define the json transition structure



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
	//choose file
	FileName := "jsons/Music--2018-08-10T09_20_31.449Z.json"

	//get json file in the form of json numbers
	jsonFile, _ := ioutil.ReadFile(FileName)

	//convert json numbers to a map, where key is string and value is a map with key string and value interface (not a fixed data structure)
	var data map[string][]map[string]interface{}
	json.Unmarshal(jsonFile, &data)

	fmt.Println(data["nodes"][1]["position"])

	//check to see if the interface type can be set to a particular correct data structure
	//var coord map[string]float64
	//coord = data["nodes"][1]["position"]

	//store all the nodes in a Graph object
	//graph := NewGraph()

}
