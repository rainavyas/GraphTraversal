import json
from pprint import pprint
import Tkinter, tkFileDialog
import os
from help import Stack

#Write a Riemann test to a file	
def write_a_test(s,m,f):

	domain=conditions['domains'] 
	intent=conditions['intents']
	k=conditions.keys()
	k.remove('domains')
	k.remove('intents')
	file.write('\tDescribe("User says ", func() { \n' '\t\tIt("Should say ", func() { \n' '\t\t\tr.Sensors.Mic.SendVT(micSensor.WAKEUP) \n'+'\t\t\tr.WaitFor(olly).ToBe(riemann.Listening()) \n'+'\t\t\tr.Sensors.Mic.SendNlpResult( \n')
	file.write('\t\t\t\tmicSensor.Domains("'+domain+'"),\n')
	file.write('\t\t\t\tmicSensor.Intents("'+intent+'"),\n')
	for entity in k:
		ent=entity.split(':')[-1]
		val=raw_input("Enter value for entity "+entity+" ")
		file.write('\t\t\t\tmicSensor.Entities("'+ent+'","'+val+'"),\n)\n')
	file.write('\t\t\tr.Match( \n'
					'\t\t\t\triemann.Group( \n'
						'\t\t\t\t\ttts.Matcher(""), \n'
					'\t\t\t\t), \n'
				'\t\t\t) \n'
			'\t\t}) \n'
		'\t}) \n')


 # Find paths
def dfs_iterative(graph, start, path):
	global paths, iteration, exit
	new_path=path[:]
	new_path.append(start.id)

	if check_iterator(start):
		# if exit:
		# 	exit=False
		# 	return new_path
		# else:
		if iteration==99:
			iteration=int(raw_input("How many cycles? (Input a positive number)")) #some iterators are in the first path out of the particular condition, if this is not alowed to loop multiple paths can be lost
		elif iteration==0: #last cycle
			iteration=99 #reset counter
			return new_path
			exit=True
		else:
			iteration=iteration-1
	if len(start.children.values())==0: #reached final node
		paths.append(new_path)
		print("fin")
		return new_path
	for n in start.children.keys(): #otherwise go through each of next nodes one by one
		neighbor=graph.nodes[start.children[n]]
		p=dfs_iterative(graph, neighbor, new_path)

def check_iterator(node):
	try:
		if node.actions[0].command=="GET_NEXT":
			return True
		else:
			return False
	except IndexError:
		return False
#Check if the condition will continue iteration
#NEED TO FIX THIS
def check_iteration(condition):
	try:
		if condition.pattern["length"]=="+":
			return True
		else:
			return False
	except KeyError:
		return False

#specify json to parse
root = Tkinter.Tk()
file = tkFileDialog.askopenfile(parent=root,initialdir="/home/andy/asr-demos/record-and-align/time_profile",mode='rb',title='Choose a file')
#with open('lights--2018-08-03T08_59_04.846Z.json') as f:
data = json.load(file)
file.close()


class Node:
	def __init__(self):
		self.id=None
		self.type=None
		self.label=None
		self.conditions = {}
		self.actions =[]
		self.children={}

	def add_connection(self, node, condition):
		self.children[condition] = node
	def dump(self):
		print("Id: "+self.id+" Type:"+self.type+" Label: "+self.label)

class Condition:
	def __init__(self):
		self.id=None
		self.type=None
		self.pattern={}

class Action:
	def __init__(self):
		self.type=None
		self.command=None
		self.parameters={}

class Graph:
	def __init__(self):
		self.nodes = {}
		self.current_node = 0

	def connect_nodes(self, i, j):
		node1 = self.nodes[i]
		node2 = self.nodes[j]
		node1.add_connection(node2)
	def dump(self):
		for n in self.nodes.values():
			print("Id: "+str(n.id)+" Type: "+n.type+" Label: "+n.label)
			for a in n.actions:
				print("Action type: "+a.type+" Command: ")
graph=Graph()

#Store all the nodes to the Graph
for i in range(len(data["nodes"])):
	node=Node()
	d=data["nodes"][i]
	node.id=data["nodes"][i]['id']
	node.type=data["nodes"][i]['type']
	node.label=data["nodes"][i]['label']
	if "conditions" in data["nodes"][i].keys():
		for cond in data["nodes"][i]['conditions']:
			c=Condition()
			c.id=cond["id"]
			c.type=cond["type"]
			for k in cond["pattern"]:
				c.pattern[k["key"]]=k["value"]
			node.conditions[c.id]=c
	if "actions" in data["nodes"][i].keys():
		for act in data["nodes"][i]['actions']:
			a=Action()
			a.type=act["type"]
			a.command=act["command"]
			for k in act["parameters"]:
				a.parameters[k["key"]]=k["value"]
			node.actions.append(a)
	graph.nodes[node.id]=node


#Add connections to graph
for i in data["transitions"]:
	start=i["from"]["node"]
	try:
		c=i["from"]["condition"]
	except KeyError:
		c=0
	end=i["to"]["node"]
	graph.nodes[start].add_connection(end, c)

#graph.dump()
exit=False
iteration=99
paths=[]
dfs_iterative(graph, graph.nodes[1],[])
print(paths)

# for p in paths:
# 	start=graph.nodes[p[0]]
# 	mid=graph.nodes[p[1]]
# 	finish=graph.nodes[p[-1]]
# 	write_a_test(start, mid, finish)
#open file to write tests in
file = open('testfile.go','w')


