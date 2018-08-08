	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("timer"),
				micSensor.Intents("timer:set"),
				micSensor.Entities("duration","+"),
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
				micSensor.Domains("timer"),
				micSensor.Intents("timer:set"),
				micSensor.Entities("time","+"),
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
				micSensor.Domains("timer"),
				micSensor.Intents("timer:set"),
				micSensor.Entities("timer_name","+"),
)
				micSensor.Entities("duration","+"),
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
				micSensor.Domains("timer:timer"),
				micSensor.Intents("timer:set"),
				micSensor.Entities("timer_name","+"),
)
				micSensor.Entities("time","+"),
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
				micSensor.Domains("timer"),
				micSensor.Intents("timer:query_get_list"),
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
				micSensor.Domains("timer"),
				micSensor.Intents("timer:query_get_list"),
				micSensor.Entities("timer_name","+"),
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
				micSensor.Domains("timer"),
				micSensor.Intents("timer:query_get_list"),
				micSensor.Entities("duration","+"),
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
				micSensor.Domains("timer"),
				micSensor.Intents("timer:query_get_list"),
				micSensor.Entities("time","+"),
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
				micSensor.Domains("timer"),
				micSensor.Intents("timer:query_get_list"),
				micSensor.Entities("timer_name","+"),
)
				micSensor.Entities("duration","+"),
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
				micSensor.Domains("timer"),
				micSensor.Intents("timer:query_get_list"),
				micSensor.Entities("method","list"),
)
				micSensor.Entities("time","+"),
)
				micSensor.Entities("timer_name","+"),
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
				micSensor.Domains("timer"),
				micSensor.Intents("timer:delete"),
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
				micSensor.Domains("timer"),
				micSensor.Intents("timer:delete"),
				micSensor.Entities("timer_name","+"),
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
				micSensor.Domains("timer"),
				micSensor.Intents("timer:delete"),
				micSensor.Entities("duration","+"),
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
				micSensor.Domains("timer"),
				micSensor.Intents("timer:delete"),
				micSensor.Entities("time","+"),
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
				micSensor.Domains("timer"),
				micSensor.Intents("timer:delete"),
				micSensor.Entities("timer_name","+"),
)
				micSensor.Entities("duration","+"),
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
				micSensor.Domains("timer"),
				micSensor.Intents("timer:delete"),
				micSensor.Entities("timer_name","+"),
)
				micSensor.Entities("time","+"),
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
				micSensor.Domains("timer"),
				micSensor.Intents("timer:set"),
			r.Match( 
				riemann.Group( 
					tts.Matcher("turning the lights on"), 
					lights.Matcher("turn_on"), 
				), 
			) 
		}) 
	}) 
