package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//define some useful global variables
var cas string
var iteration = 99
var paths [][][]int

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
func (n Node) add_child(childID, conditionID int) Node {
	n.Children[conditionID] = childID
	return n

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

//
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

//Depth First Search Algorithm to find paths
func DfsIterative(graph Graph, start Node, path [][]int) {
	if (CheckIterator(start)) && (len(start.Children) != 0) {
		//the current node is an iterator
	}

	newPath := path
	if len(start.Children) == 0 || start.Id == 0 {
		//reached final node or 0th node
		if start.Id != 0 {
			tempPath := make([]int, 2)
			tempPath[0] = start.Id
			tempPath[1] = 0
			newPath = append(newPath, tempPath)
		}
		paths = append(paths, newPath)
		return
	}

	//go through each of the conditions of the current node
	for k, v := range start.Children {
		tempPath := make([]int, 2)
		tempPath[0] = start.Id
		tempPath[1] = k
		//tempPath = append(tempPath, k)
		newPath = append(newPath, tempPath)
		neighbour := graph.Nodes[v]
		DfsIterative(graph, neighbour, newPath)

	}

}

//function to check if node is iterator
func CheckIterator(node Node) bool {
	if len(node.Actions) > 0 {
		if node.Actions[0].Command == "GET_NEXT" {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func main() {
	//choose file
	FileName := "jsons/Music--2018-08-10T09_20_31.449Z.json"
	//FileName := "jsons/Who-Am I--2018-08-10T09_22_10.343Z.json"
	//FileName := "jsons/Wifi--2018-08-10T09_22_13.471Z.json"

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

	graph := NewGraph()
	for i := 0; i < len(JsonStructs.Nodes); i++ {
		//populate Node structure
		CurrentNode := NewNode()
		CurrentNode.Id = JsonStructs.Nodes[i].Id
		CurrentNode.Typ = JsonStructs.Nodes[i].Typ
		CurrentNode.Label = JsonStructs.Nodes[i].Label

		//loop through the conditions
		for j := 0; j < len(JsonStructs.Nodes[i].Conditions); j++ {
			CurrentCond := NewCondtion()
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
			CurrentAction := NewAction()
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
	for _, transition := range JsonStructs.Transitions {

		nodeFromIndex := transition.From.Node

		tempNode := graph.Nodes[nodeFromIndex]

		tempNode = tempNode.add_child(
			transition.To.Node,
			transition.From.Condition,
		)

		graph.Nodes[transition.From.Node] = tempNode

	}

	// for k := range graph.Nodes {
	// 	fmt.Println("the final nodes list")
	// 	fmt.Println(k)
	// 	fmt.Println(graph.Nodes[k].Children)
	// }

	//call dfs iterative
	var path [][]int
	DfsIterative(graph, graph.Nodes[1], path)
	fmt.Print(paths)

	//fmt.Println(len(graph.Nodes[7].Children))

	/*testMap := make(map[int]int)
	testMap[0] = 1
	testMap[1] = 2
	testMap[2] = 3

	fmt.Println(len(testMap))*/

}
