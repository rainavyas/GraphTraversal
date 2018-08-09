import json
from pprint import pprint

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
	
def write_a_test(conditions,extra):
	domain=conditions['domains'] 
	intent=conditions['intents']
	k=conditions.keys()
	k.remove('domains')
	k.remove('intents')
	file.write('\tDescribe("User says ", func() { \n' '\t\tIt("Should say ", func() { \n' '\t\t\tr.Sensors.Mic.SendVT(micSensor.WAKEUP) \n'+'\t\t\tr.WaitFor(olly).ToBe(riemann.Listening()) \n'+'\t\t\tr.Sensors.Mic.SendNlpResult( \n')
	file.write('\t\t\t\tmicSensor.Domains("'+domain+'"),\n')
	file.write('\t\t\t\tmicSensor.Intents("'+intent+'"),\n')
	for entity in k:
		val=conditions[entity]
		ent=entity.split(':')[-1]
		file.write('\t\t\t\tmicSensor.Entities("'+ent+'","'+val+'"),\n)\n')
	file.write('\t\t\tr.Match( \n'
					'\t\t\t\triemann.Group( \n'
						'\t\t\t\t\ttts.Matcher("'+extra+'"), \n'
					'\t\t\t\t), \n'
				'\t\t\t) \n'
			'\t\t}) \n'
		'\t}) \n')
 
	
#specify json to parse
with open('lights--2018-08-03T08_59_04.846Z.json') as f:
    data = json.load(f)

#store the domains, intents and entities
parsed=[]

#Go through each condition in node 1 and capture domain, intent and entity in dictionaries in a list
for j in range(len(data["nodes"][1]['conditions'])):
	cond=data["nodes"][1]['conditions'][j]['id']
	d={}
	for i in range(len(data["nodes"][1]['conditions'][j]['pattern'])):
		try:
			d[str(data["nodes"][1]['conditions'][j]['pattern'][i]['key'])]=str(data["nodes"][1]['conditions'][j]['pattern'][i]['value'])
		except IndexError:
			pass
	parsed.append(d)

#open file to write tests in
file = open('testfile.go','w')

#find one path for each starting conditions
tts=[]
for i in range(len(data["nodes"][1]['conditions'])):
	path=[]
	fin, path=find_next(data["nodes"][1],path)
	print("End:", fin)
	x=(id_to_idx(path[-2]))
	tts.append(data["nodes"][x]["actions"][0]["parameters"][0]["value"])


#write tests
for i,j in zip(parsed, tts):
	write_a_test(i,j)

