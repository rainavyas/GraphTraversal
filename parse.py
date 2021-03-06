import json
from pprint import pprint
import Tkinter, tkFileDialog
import os
from help import Stack

#function to convert node id to idx in the nodes list
def id_to_idx(noid):
	for i in range(len(data["nodes"])):
		node_id=data["nodes"][i]["id"]
		if node_id==noid:
			return(i)

#function to find next node given a start node
#Output: {condition number: next_node_id}, [first_node_id, middle_node_id ... last_node_id]
def find_next(start_node,path):
	n=start_node["id"]
	print("Start:",n)
	path.append(n)
	ends={} #prepare storage for all possible next nodes
	if "conditions" not in start_node.keys() or len(start_node["conditions"])<=1: #node with 0 or 1 condition
		i=0
		try:
			while data["transitions"][i]['from']['node']!=n:
				i+=1
			ends['0']=data["transitions"][i]['to']['node'] #next node found, find next next node
			x=id_to_idx(ends['0'])
			ends,path=find_next(data["nodes"][x],path)
		except IndexError: #the node has no outgoing connection i.e. is the last node
			ends['0']=n
			return ends, path
	else: 
		for i in range(len(data["transitions"])): #find all possible next nodes
			if data["transitions"][i]['from']['node']==n:
				cond=data["transitions"][i]['from']['condition']
				next=data["transitions"][i]['to']['node']
				ends[cond]=next
		print(ends)
		choice=int(raw_input("Which condition?: ")) #choose condition to take for next step
		x=ends[choice]
		x=id_to_idx(x)
		ends,path=find_next(data["nodes"][x],path)
	return ends,path

#Write a Riemann test to a file	
def write_a_test(conditions):
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
	print start.children
	global paths, iteration, exit
	print "Iteration"+str(iteration)
	new_path=path[:]
	new_path.append(start.id)

	if len(start.children.values())==0: #reached final node
		paths.append(new_path)
		print("fin")
		return new_path
	for n in start.children.keys(): #otherwise go through each of next nodes one by one
		neighbor=graph.nodes[start.children[n]]
		try:
			if check_iteration_end(start.conditions[n]):
				if exit:
					exit=False
					neighbor=graph.nodes[start.children.values()[1]]
				if iteration==99:
					iteration=int(raw_input("How many cycles?"))
				elif iteration==0: #last cycle
						iteration=99 #reset counter
						exit=True
				else:
					iteration=iteration-1
		except KeyError:
			pass
		p=dfs_iterative(graph, neighbor, new_path)

def check_iterator(node):
	try:
		print node.actions[0].type
		iterator=node.actions[0].type=="ITERATOR"
		return iterator
	except IndexError:
		return False

def check_iteration_end(condition):
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


#store the domains, intents and entities
parsed=[]



class Node:
	def __init__(self):
		self.id=None
		self.type=None
		self.label=None
		self.conditions = {}
		self.actions =[]
		self.children={}

	def add_connection(self, node, condition):
		#num_conditions = len(self.conditions.items())
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
		#node.tts=(data["nodes"][x]["actions"][0]["parameters"][0]["value"])
	graph.nodes[node.id]=node



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
#Go through each condition in node 1 and capture domain, intent and entity in dictionaries in a list
# for j in range(len(data["nodes"][1]['conditions'])):
# 	cond=data["nodes"][1]['conditions'][j]['id']
# 	d={}
# 	for i in range(len(data["nodes"][1]['conditions'][j]['pattern'])):
# 		try:
# 			d[str(data["nodes"][1]['conditions'][j]['pattern'][i]['key'])]=str(data["nodes"][1]['conditions'][j]['pattern'][i]['value'])
# 		except IndexError:
# 			pass
# 	parsed.append(d)

#open file to write tests in
file = open('testfile.go','w')

# #find one path for each starting conditions
# for i in range(len(graph.nodes[1].conditions.items)):
# 	next=graph.nodes[1].children[i]
# 	path=[]
# 	fin.id, path=find_path(graph.nodes[1].conditions,path)
# 	print("End:", fin.id)
# 	x=(id_to_idx(path[-2]))


# #write tests
# for i,j in zip(parsed, tts):
# 	write_a_test(i,j)

