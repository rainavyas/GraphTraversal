	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("time"),
				micSensor.Intents("time:query"),
)			r.Match( 
				riemann.Group( 
					tts.Matcher("it is [time]"), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("time"),
				micSensor.Intents("time:query"),
				micSensor.Entities("location","val"),
)
)			r.Match( 
				riemann.Group( 
					tts.Matcher("it is [time] in [location]"), 
				), 
			) 
		}) 
	}) 
