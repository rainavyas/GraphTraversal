import json
from pprint import pprint
import Tkinter, tkFileDialog
import os
from help import Stack, Queue
import collections 

#Write a Riemann test to a file	in GO syntax
def write_a_test(s,c,f):
	#open file to write tests in
	domain=s.conditions[c].pattern['domains'] 
	intent=s.conditions[c].pattern['intents']
	k=s.conditions[c].pattern.keys()
	try:
		tts=f.actions[0].parameters["sentence"]
	except KeyError:
		tts="No TTS"
	k.remove('domains')
	k.remove('intents')
	file.write('\tDescribe("User says ", func() { \n' '\t\tIt("Should say ", func() { \n' '\t\t\tr.Sensors.Mic.SendVT(micSensor.WAKEUP) \n'+'\t\t\tr.WaitFor(olly).ToBe(riemann.Listening()) \n'+'\t\t\tr.Sensors.Mic.SendNlpResult( \n')
	file.write('\t\t\t\tmicSensor.Domains("'+domain+'"),\n')
	file.write('\t\t\t\tmicSensor.Intents("'+intent+'"),\n')
	for entity in k:
		ent=entity.split(':')[-1]
		val="val"
		#raw_input("Enter value for entity "+entity+" ")
		file.write('\t\t\t\tmicSensor.Entities("'+ent+'","'+val+'"),\n)\n')
	file.write(')')
	file.write('\t\t\tr.Match( \n'
					'\t\t\t\triemann.Group( \n'
						'\t\t\t\t\ttts.Matcher("'+tts+'"), \n'
					'\t\t\t\t), \n'
				'\t\t\t) \n'
			'\t\t}) \n'
		'\t}) \n')


 # Find paths
def dfs_iterative(graph, start, path):
	global paths, iteration, case
	if (check_iterator(start) and len(start.children.values()) != 0):
		if iteration==99:
			"""first time reached iterator"""
			iteration=int(raw_input("How many cycles? (Input a positive number)"))
			
			if start.children[0] in path:
				"""node after iterator has come up before"""
				case = "A"
				if iteration == 0:
					"""quit this loop"""
					iteration = 99
					return
				else:
					iteration = iteration-1 #necessary to get correct cycles for case A
			else:
				case = "B" 

		elif iteration==0: #last cycle
			iteration=99 #reset counter
			return  #prematurely end this loop
		else:
			iteration=iteration-1
	new_path=path[:]
	if len(start.children.values())==0 or start.id==0: #reached final node
		if start.id!=0:
			new_path.append(start.id,0)
		paths.append(new_path)
		return
	new_path.append((0,0))
	for n in start.children.keys(): #otherwise go through each of next nodes one by one
		new_path[-1]=(start.id,n)
		neighbor=graph.nodes[start.children[n]]
		dfs_iterative(graph, neighbor,new_path)

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
root.update()
root.destroy()
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
			node.children[c.id] = 0
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
	graph.nodes[start].children[c]=end

#define useful global variables
case = "C"
iteration=99
paths=[]
total_its =0

#Store each path as a list, where each element is [nodeID, conditionID to next node]

dfs_iterative(graph, graph.nodes[1],[])
print(paths)

#print paths[0][1]


file = open('testfile.go','w')
for p in paths:
	#print p
	start=graph.nodes[p[0][0]]
	c=p[0][1]
	#print start.id
	finish=graph.nodes[p[-2][0]]
	#print finish.id
	write_a_test(start,c, finish)
file.close()



