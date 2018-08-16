	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("weather"),
				micSensor.Intents("weather:query_condition"),
)			r.Match( 
				riemann.Group( 
					tts.Matcher("The weather is [weather:description] and the temperature is [weather:temperature_value] degrees"), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("weather"),
				micSensor.Intents("weather:query_condition"),
)			r.Match( 
				riemann.Group( 
					tts.Matcher("Sorry, I can not fetch the weather right now"), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("weather"),
				micSensor.Intents("weather:query_condition"),
				micSensor.Entities("location","val"),
)
)			r.Match( 
				riemann.Group( 
					tts.Matcher("In [weather:location] the weather is [weather:description] and the temperature is [weather:temperature_value] degrees"), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("weather"),
				micSensor.Intents("weather:query_condition"),
				micSensor.Entities("location","val"),
)
)			r.Match( 
				riemann.Group( 
					tts.Matcher("Sorry, I can not fetch the weather right now"), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("weather"),
				micSensor.Intents("weather:query_condition"),
				micSensor.Entities("time_range","val"),
)
)			r.Match( 
				riemann.Group( 
					tts.Matcher("In [weather:location] the weather [weather:weekday] is [weather:description] with highs of [weather:max_temperature_value] and lows of [weather:min_temperature_value]."), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("weather"),
				micSensor.Intents("weather:query_condition"),
				micSensor.Entities("time_range","val"),
)
)			r.Match( 
				riemann.Group( 
					tts.Matcher("Sorry, I can not fetch the weather right now"), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("weather"),
				micSensor.Intents("weather:query_condition"),
				micSensor.Entities("time_range","val"),
)
				micSensor.Entities("location","val"),
)
)			r.Match( 
				riemann.Group( 
					tts.Matcher("In [weather:location] the weather [weather:weekday] is [weather:description] with highs of [weather:max_temperature_value] and lows of [weather:min_temperature_value]."), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("weather"),
				micSensor.Intents("weather:query_condition"),
				micSensor.Entities("time_range","val"),
)
				micSensor.Entities("location","val"),
)
)			r.Match( 
				riemann.Group( 
					tts.Matcher("Sorry, I can not fetch the weather right now"), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("weather"),
				micSensor.Intents("weather:query_condition"),
				micSensor.Entities("time_range","val"),
)
)			r.Match( 
				riemann.Group( 
					tts.Matcher("In [weather:location] the weather [weather:weekday] is [weather:description] with highs of [weather:max_temperature_value] and lows of [weather:min_temperature_value]."), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("weather"),
				micSensor.Intents("weather:query_condition"),
				micSensor.Entities("time_range","val"),
)
)			r.Match( 
				riemann.Group( 
					tts.Matcher("Sorry, I can not fetch the weather right now"), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("weather"),
				micSensor.Intents("weather:query_condition"),
				micSensor.Entities("time_range","val"),
)
				micSensor.Entities("location","val"),
)
)			r.Match( 
				riemann.Group( 
					tts.Matcher("In [weather:location] the weather [weather:weekday] is [weather:description] with highs of [weather:max_temperature_value] and lows of [weather:min_temperature_value]."), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("weather"),
				micSensor.Intents("weather:query_condition"),
				micSensor.Entities("time_range","val"),
)
				micSensor.Entities("location","val"),
)
)			r.Match( 
				riemann.Group( 
					tts.Matcher("Sorry, I can not fetch the weather right now"), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("weather"),
				micSensor.Intents("weather:query_condition"),
				micSensor.Entities("time_range","val"),
)
)			r.Match( 
				riemann.Group( 
					tts.Matcher("Sorry, I can not fetch the weather right now"), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("weather"),
				micSensor.Intents("weather:query_condition"),
				micSensor.Entities("time_range","val"),
)
				micSensor.Entities("location","val"),
)
)			r.Match( 
				riemann.Group( 
					tts.Matcher("Sorry, I can not fetch the weather right now"), 
				), 
			) 
		}) 
	}) 
