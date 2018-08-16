import json
from pprint import pprint
import Tkinter, tkFileDialog
import os
from help import Stack


 

	
def check_iterator(node):
	"""is the current node an iterator"""
	try:
		if node.actions[0].command=="GET_NEXT":
			return True
		else:
			return False
	except IndexError:
		return False


def first_time(nodeID, path):
	"check if node has come up before"
	if nodeID in path:
		return False
	else:
		return True




def recursive_path_finder(graph, curr_node, path):	

	global paths, its, case, val 
	#let the new_path hold the path list
	new_path=path[:]


	if (check_iterator(curr_node) and len(curr_node.children.values()) != 0):
		""" iterator node reached and not the last node"""
		if first_time(curr_node.id, new_path):
			its=1
			exit = False
			val = int(raw_input("How many cycles?"))

			if curr_node.children[0] in new_path:
				"""node after iterator has come up before"""
				case = "A"
			else:
				case = "B"
				val = val +1 #necessary for edge conditions of case B
		else:
			its=its+1

		
		if its > val:
			"""need to terminate path"""
			return
				#while (new_path[-1] != curr_node.children[0]):
				#	remove last element in path
				#	del new_path[-1]
				#del new_path[-1]
				#exit = True
				#recursive_path_finder(graph, graph.nodes[curr_node.children[0]], new_path)
			#else:
			#	new_path.append(curr_node.id)
			#	recursive_path_finder(graph, graph.nodes[curr_node.children[0]], new_path)


		
	#add current node
	new_path.append(curr_node.id)

	if len(curr_node.children.values())==0:
		"""on final node"""
		if new_path not in paths:
			#no duplicate paths
			paths.append(new_path)
		

	else:
		"""go through each condition in turn"""
		for cond in curr_node.children.keys():
			new_node = graph.nodes[curr_node.children[cond]]
			recursive_path_finder(graph, new_node, new_path)






	
	

	





#create GUI to select json to parse
root = Tkinter.Tk()
file = tkFileDialog.askopenfile(parent=root,initialdir="/home/andy/asr-demos/record-and-align/time_profile",mode='rb',title='Choose a file')
data = json.load(file)
file.close()


class Node:
	""" defining class for each node"""
	def __init__(self):
		self.id=None
		self.type=None
		self.label=None
		self.conditions = {} #key: condition ID , value: condition
		self.actions =[] #every item is an Action object
		self.children={} #key: condition , value: child node id

	def add_connection(self, nodeID, condition):
		"""add connecting nodes via a dictionary with condition as the key"""
		self.children[condition] = nodeID 

	def dump(self):
		"""output the attribute values of node"""
		print("Id: "+self.id+" Type:"+self.type+" Label: "+self.label)

class Condition:
	def __init__(self):
		self.id=None
		self.type=None
		self.pattern={} #key: key ; value: value

class Action:
	def __init__(self):
		self.type=None
		self.command=None
		self.parameters={} #key: key ; value: value

class Graph:
	def __init__(self):
		self.nodes = {} #key: node ID; value: node object
		self.current_node = 0


	def dump(self):
		for n in self.nodes.values():
			print("Id: "+str(n.id)+" Type: "+n.type+" Label: "+n.label)
			for a in n.actions:
				print("Action type: "+a.type+" Command: ")


#Create graph object				
graph=Graph()

#Store all the nodes to the Graph
#loop through list of nodes in json file
for i in range(len(data["nodes"])):
	#create Node object to hold current node data
	node=Node()

	#fill in attribute values of current node
	node.id=data["nodes"][i]['id']
	node.type=data["nodes"][i]['type']
	node.label=data["nodes"][i]['label']

	#add conditions to current node if they exist
	if "conditions" in data["nodes"][i].keys():
		for cond in data["nodes"][i]['conditions']:
			#create a current condition object, c
			c=Condition()
			c.id=cond["id"]
			c.type=cond["type"]
			#loop through all the patterns
			for pat in cond["pattern"]:
				c.pattern[pat["key"]]=pat["value"]
			#assign the completed current condition object to dictionary of conditions of the current node	
			node.conditions[c.id]=c
	#add actions to current node if they exist		
	if "actions" in data["nodes"][i].keys():
		for act in data["nodes"][i]['actions']:
			#create a current action object, a
			a=Action()
			a.type=act["type"]
			a.command=act["command"]
			for par in act["parameters"]:
				a.parameters[par["key"]]=par["value"]
			#append the current action object to the list of actions of the node	
			node.actions.append(a)

	#add current node to dictionary of nodes in graph object
	graph.nodes[node.id]=node


#Add connections to graph
#loop through list of transitions in JSON file
for trans in data["transitions"]:
	#get start node ID
	start=trans["from"]["node"]
	#check to see if condition exists
	try:
		c=trans["from"]["condition"]
	except KeyError:
		c=0
	#get end node ID
	end=trans["to"]["node"]
	#add transition line to correct node
	graph.nodes[start].add_connection(end, c)


#global variables
exit=False
its=0
paths=[] #going to be list of all possible paths
nxt=0
case = "C"
val = 0

#call function to find all possible paths in graph- starting from node 1 as node 0 is irrelevant
#dfs_iterative(graph, graph.nodes[1],[])
recursive_path_finder(graph, graph.nodes[1],[])

print(paths)







# for p in paths:
# 	start=graph.nodes[p[0]]
# 	mid=graph.nodes[p[1]]
# 	finish=graph.nodes[p[-1]]
# 	write_a_test(start, mid, finish)
#open file to write tests in
#file = open('testfile.go','w')


