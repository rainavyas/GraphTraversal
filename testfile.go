	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_on"),
			r.Match( 
				riemann.Group( 
					tts.Matcher("turning the lights on"), 
					lights.Matcher("turn_on"), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_on"),
				micSensor.Entities("house_place","+"),
)
			r.Match( 
				riemann.Group( 
					tts.Matcher("turning the lights on"), 
					lights.Matcher("turn_on"), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_on"),
				micSensor.Entities("house_place","+"),
)
				micSensor.Entities("color","+"),
)
			r.Match( 
				riemann.Group( 
					tts.Matcher("turning the lights on"), 
					lights.Matcher("turn_on"), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_up"),
			r.Match( 
				riemann.Group( 
					tts.Matcher("turning the lights on"), 
					lights.Matcher("turn_on"), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_up"),
				micSensor.Entities("house_place","+"),
)
			r.Match( 
				riemann.Group( 
					tts.Matcher("turning the lights on"), 
					lights.Matcher("turn_on"), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_down"),
			r.Match( 
				riemann.Group( 
					tts.Matcher("turning the lights on"), 
					lights.Matcher("turn_on"), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_down"),
				micSensor.Entities("house_place","+"),
)
			r.Match( 
				riemann.Group( 
					tts.Matcher("turning the lights on"), 
					lights.Matcher("turn_on"), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:set_specific"),
				micSensor.Entities("house_place","+"),
)
				micSensor.Entities("color","+"),
)
			r.Match( 
				riemann.Group( 
					tts.Matcher("turning the lights on"), 
					lights.Matcher("turn_on"), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:set_specific"),
				micSensor.Entities("color","+"),
)
			r.Match( 
				riemann.Group( 
					tts.Matcher("turning the lights on"), 
					lights.Matcher("turn_on"), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_off"),
			r.Match( 
				riemann.Group( 
					tts.Matcher("turning the lights on"), 
					lights.Matcher("turn_on"), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_off"),
				micSensor.Entities("house_place","+"),
)
			r.Match( 
				riemann.Group( 
					tts.Matcher("turning the lights on"), 
					lights.Matcher("turn_on"), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:set_specific"),
			r.Match( 
				riemann.Group( 
					tts.Matcher("turning the lights on"), 
					lights.Matcher("turn_on"), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:set_specific"),
				micSensor.Entities("house_place","+"),
)
			r.Match( 
				riemann.Group( 
					tts.Matcher("turning the lights on"), 
					lights.Matcher("turn_on"), 
				), 
			) 
		}) 
	}) 
