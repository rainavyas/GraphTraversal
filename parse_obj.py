import json
from pprint import pprint
import Tkinter as tk
import tkFileDialog
import os
from help import Stack, Queue
import shutil

class GUI:
	def __init__(self, master,text):
		self.master = master
		master.title("Complete the scenario")
		master.configure(background='lavender blush')

		self.label1=tk.Label(master, text="User says:").grid(row=0, column=0, sticky=tk.W)
		self.label2=tk.Label(master, text="Olly should say:").grid(row=1, column=0, sticky=tk.W)

		#Strings to be written to file
		self.user= tk.StringVar() #User says
		self.olly= tk.StringVar() #Olly should say

		self.e1 = tk.Entry(master,textvariable=self.user)
		self.e2 = tk.Entry(master,textvariable=self.olly)
		self.e1.grid(row=0, column=1,sticky=tk.W+tk.E)
		self.e2.grid(row=1, column=1, sticky=tk.W+tk.E)

		self.save=tk.Button(master, text='Save', command=master.destroy, width=80,bg = 'LightPink1').grid(row=4, column=0,columnspan=3)

		self.label3=tk.Label(master, text="Current test").grid(row=2, column=0, columnspan=3)
		self.S = tk.Scrollbar(master)
		self.S.grid(row=3, column=3,sticky=tk.N+tk.S)
		self.T = tk.Text(master,height = 10, width=80)
		self.S.config(command=self.T.yview)
		self.T.grid(row=3, columnspan=3)
		self.T.insert(tk.END, text)
		self.T.config(yscrollcommand=self.S.set)
		self.T.configure(state='disabled')
		
#Write a Riemann test to a file	in GO syntax
def write_a_test(elist):
	fulltext="" #The full text of the current scenario

	tts="" #The text said by olly

	ents=[] #entities (speech recognised)
	vals=[] #values of entities e.g. Tokyo for enitity weather:loacation

	matcher="" #The name of the .Matcher needed (i.e what service is used)
	commands={} #Action.command
	params={} #Additional information for action nodes

	for e in elist: #check what events are in the path
		if e.type=="TTS": #TTS node
			tts=e.parameters["sentence"] #put all the tts in one string*
			if tts!='':
				fulltext+=('\t\t\t\t\ttts.Matcher("'+tts+'"), \n') 
		elif e.type==5:#Speech Recognised event
			keys=e.pattern.keys() #Get the NLP information
			for k in keys: 
				if k=='domains':
					domain=e.pattern["domains"]
				elif k=='intents':
					intent=e.pattern["intents"]
				else: #enitity
					ents.append(k.split(':')[-1])#store all entities, get the latter part of entity e.g. iot:location--->location
					vals.append("val")
					#raw_input("Enter value for entity "+entity+" ")
			fulltext+=('\t\t\tr.Sensors.Mic.SendNlpResult( \n')
			fulltext+=('\t\t\t\tmicSensor.Domains("'+domain+'"),\n')
			fulltext+=('\t\t\t\tmicSensor.Intents("'+intent+'"),\n')
			if len(ents)>0:
				fulltext+=('\t\t\t\tmicSensor.Entities(\n')
				for ent, val in zip(ents, vals):
					fulltext+=('\t\t\t\t\t"'+ent+'","'+val+'"\n')
				fulltext+=('\t\t\t\t)\n')
			fulltext+=('\t\t\t)\n')
			fulltext+=('\t\t\tr.Match( \n'
							'\t\t\t\triemann.Group( \n')
		elif type(e.type)!=int: #Any other action node
			action_name=e.type.lower().split('.')[-1] #Split to ensure the action name is a single word e.g iot.LIGHTS--->lights
			commands[action_name]=e.command.lower() #Save the command under the action type
			params[action_name]=e.parameters.keys() #Save the list of action parameter names under the action type
			#Write Action Macthers
			if len(commands.keys())>0:
				for matcher, command in zip(commands.keys(),commands.values()):
					if matcher=="iterator": #skip iterators
						continue
					fulltext+=('\t\t\t\t\t'+matcher+'.Matcher("'+command+'", \n')
					for param in params[matcher]:
						fulltext+=('\t\t\t\t\t\t\thass.Entity("'+param+'","val")\n') #hass.Entity needs to be replaced with appropriate name
					fulltext+=('\t\t\t\t\t)\n')
		else:
			pass
	fulltext+=('\t\t\t\t), \n'
					'\t\t\t) \n'
				'\t\t}) \n'
			'\t}) \n')

	#Open GUI to complete scenario
	master = tk.Tk()
	my_gui = GUI(master,fulltext)
	master.mainloop()

	#Retrieve input
	user=my_gui.user.get()
	olly=my_gui.olly.get()

	#Add the first two lines with user readable scenario summary at the start
	fulltext=('\tDescribe("User says '+user+'", func() { \n' '\t\tIt("Should say '+olly+'", func() { \n' '\t\t\tr.Sensors.Mic.SendVT(micSensor.WAKEUP) \n'+'\t\t\tr.WaitFor(olly).ToBe(riemann.Listening()) \n')+fulltext
	#write scenario to file
	file.write(fulltext)


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
			new_path.append((start.id,0))
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
def check_iteration(condition):
	try:
		if condition.pattern["length"]=="+":
			return True
		else:
			return False
	except KeyError:
		return False

#specify json to parse
root = tk.Tk()
file = tkFileDialog.askopenfile(parent=root,initialdir='/jsons',mode='rb',title='Choose a file')
name=str(file).split('/')[-1].split('-')[0]
root.update()
root.destroy()
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
#print(paths)

#Copy the template
shutil.copy('template.go', 'tests/'+name+'.go')

#Get all the conditions and actions and write tests for each path to one file

file = open('tests/'+name+'.go','a')
for p in paths:
	events=[]
	for stop in p:
		loc=graph.nodes[stop[0]]
		try:
			events.append(loc.conditions[stop[1]]) #Get the one chosen Condition in the node
		except:
			KeyError
		for i in loc.actions: #Get all the Actions in node
			events.append(i)

	write_a_test(events)
file.close()



