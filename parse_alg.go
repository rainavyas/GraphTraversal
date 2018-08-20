package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//define an Action structure
type Action struct {
	Typ        string
	Command    string
	Parameters map[string]string
}

//define a Condition structure
type Condition struct {
	Id      int
	Typ     int
	Pattern map[string]string //each new key,val is a new pattern
}

// define a node structure
type Node struct {
	Id         int
	Typ        string
	Label      string
	Conditions map[int]Condition
	Actions    []Action
	Children   map[int]int
}

//create method for Node class to be able to add children nodes
func (n Node) add_child(childID, conditionID int) {
	n.Children[conditionID] = childID
}

//define a Graph structure
type Graph struct {
	Nodes map[int]Node
}

//define a set of structures required for json unmarsheling

//define json overall tree
type JsonNodesAndTrans struct {
	Nodes       []JsonNode       `json:"nodes"`
	Transitions []JsonTransition `json:"transitions"`
}

//define the json node structure
type JsonNode struct {
	Id         int             `json:"id"`
	Typ        string          `json:"type"`
	Label      string          `json:"label`
	Position   JsonPosition    `json:"position"`
	Conditions []JsonCondition `json:"conditions"`
	Actions    []JsonAction    `json:"actions"`
}

//define the json position structure
type JsonPosition struct {
	X int `json:"x"`
	Y int `json:"y"`
}

//define the json condition structure
type JsonCondition struct {
	Id      int           `json:"id"`
	Typ     int           `json:"type"`
	Pattern []JsonPattern `json:"pattern"`
}

//define the json pattern structure
type JsonPattern struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

//define the json action structure
type JsonAction struct {
	Typ        string          `json:"type"`
	Command    string          `json:"command"`
	Parameters []JsonParameter `json:"parameters"`
}

//define the json paramater structure
type JsonParameter struct {
	Key   string `json:"key"`
	Value string `json"value"`
}

//define the json transition structure
type JsonTransition struct {
	From JsonFrom `json:"from"`
	To   JsonTo   `json:"to"`
}

//define the json from structure
type JsonFrom struct {
	Node      int `json:"node"`
	Condition int `json:"condition"`
}

//define the json to structure
type JsonTo struct {
	Node int `json:"node"`
}

//function to create Action object
func NewAction() Action {
	return Action{Parameters: map[string]string{}}
}

//function to create Condition object
func NewCondtion() Condition {
	return Condition{Pattern: map[string]string{}}
}

// function to create Node object
func NewNode() Node {
	return Node{Conditions: map[int]Condition{}, Children: map[int]int{}}
}

//function to create Graph object
func NewGraph() Graph {
	return Graph{Nodes: map[int]Node{}}
}

func main() {
	//choose file
	FileName := "jsons/Music--2018-08-10T09_20_31.449Z.json"
	//FileName := "Music--2018-08-10T09_20_31.449Z.json"

	//get json file in the form of json numbers
	jsonFile, err := os.Open(FileName)

	//if os.Open returns an error, handle it
	if err != nil {
		fmt.Println(err)
	}

	//defer the closing of the jsonFile so that it can be parsed later
	defer jsonFile.Close()

	//read the opened file as byte array
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}

	//initialise the json structures
	JsonStructs := JsonNodesAndTrans{}

	//umarshal the byteArray into the json structures defined
	err = json.Unmarshal(byteValue, &JsonStructs)
	if err != nil {
		fmt.Println("error reached", err)
	}

	//store all the nodes in a Graph object
	CurrentNode := NewNode()
	CurrentCond := NewCondtion()
	CurrentAction := NewAction()
	graph := NewGraph()
	for i := 0; i < len(JsonStructs.Nodes); i++ {
		//populate Node structure
		CurrentNode.Id = JsonStructs.Nodes[i].Id
		CurrentNode.Typ = JsonStructs.Nodes[i].Typ
		CurrentNode.Label = JsonStructs.Nodes[i].Label

		//loop through the conditions
		for j := 0; j < len(JsonStructs.Nodes[i].Conditions); j++ {
			CurrentCond.Id = JsonStructs.Nodes[i].Conditions[j].Id
			CurrentCond.Typ = JsonStructs.Nodes[i].Conditions[j].Typ

			//loop through the patterns
			for k := 0; k < len(JsonStructs.Nodes[i].Conditions[j].Pattern); k++ {
				CurrentCond.Pattern[JsonStructs.Nodes[i].Conditions[j].Pattern[k].Key] = JsonStructs.Nodes[i].Conditions[j].Pattern[k].Value
			}

			CurrentNode.Conditions[CurrentCond.Id] = CurrentCond
		}

		//loop through the actions
		for a := 0; a < len(JsonStructs.Nodes[i].Actions); a++ {
			CurrentAction.Command = JsonStructs.Nodes[i].Actions[a].Command
			CurrentAction.Typ = JsonStructs.Nodes[i].Actions[a].Typ

			//loop through the parameters
			for p := 0; p < len(JsonStructs.Nodes[i].Actions[a].Parameters); p++ {
				CurrentAction.Parameters[JsonStructs.Nodes[i].Actions[a].Parameters[p].Key] = JsonStructs.Nodes[i].Actions[a].Parameters[p].Value
			}

			CurrentNode.Actions = append(CurrentNode.Actions, CurrentAction)

		}

		graph.Nodes[CurrentNode.Id] = CurrentNode

	}

	//add connections between the nodes in the graph
	for t := 0; t < len(JsonStructs.Transitions); t++ {
		graph.Nodes[JsonStructs.Transitions[t].From.Node].add_child(JsonStructs.Transitions[t].To.Node, JsonStructs.Transitions[t].From.Condition)
	}

}
