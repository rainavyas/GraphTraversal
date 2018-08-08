import json
from pprint import pprint

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
		val=conditions[entity]
		ent=entity.split(':')[-1]
		file.write('\t\t\t\tmicSensor.Entities("'+ent+'","'+val+'"),\n)\n')
	file.write('\t\t\tr.Match( \n'
					'\t\t\t\triemann.Group( \n'
						'\t\t\t\t\ttts.Matcher("turning the lights on"), \n'
						'\t\t\t\t\tlights.Matcher("turn_on"), \n'
					'\t\t\t\t), \n'
				'\t\t\t) \n'
			'\t\t}) \n'
		'\t}) \n')
 
	

with open('Timers--2018-08-08T10_11_37.473Z.json') as f:
    data = json.load(f)
parsed=[]

print(len(data["nodes"][1]['conditions']))
print(len(data["nodes"][1]['conditions'][1]['pattern']))
for j in range(len(data["nodes"][1]['conditions'])):
	d={}
	for i in range(len(data["nodes"][1]['conditions'][j]['pattern'])):
		try:
			d[str(data["nodes"][1]['conditions'][j]['pattern'][i]['key'])]=str(data["nodes"][1]['conditions'][j]['pattern'][i]['value'])
		except IndexError:
			pass
	parsed.append(d)

for i in parsed:
	print(i)
	print('\n')
print(len(parsed))


file = open('testfile.go','w') 
 
for i in parsed:
	write_a_test(i)
 
 
# file.close()
# Describe("User says ''", func() {
# 		It("Should say '", func() {
# 			r.Sensors.Mic.SendVT(micSensor.WAKEUP)
# 			r.WaitFor(olly).ToBe(riemann.Listening())
# 			r.Sensors.Mic.SendNlpResult(
# 				micSensor.Domains(domain),
# 				micSensor.Intents(intent),
# 				micSensor.Entities(),
# 			)
# 			r.Match(
# 				riemann.Group(
# 					tts.Matcher("turning the lights on"),
# 					lights.Matcher("turn_on"),
# 				),
# 			)
# 		})
# 	})