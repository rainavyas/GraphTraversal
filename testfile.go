	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("timer"),
				micSensor.Intents("timer:set"),
				micSensor.Entities(
					"duration","val"
				)
)			r.Match( 
				riemann.Group( 
					timers.Matcher("timers", 
							hass.Entity("duration","val")
							hass.Entity("name","val")
							hass.Entity("time","val")
					timeout.Matcher("timeout", 
							hass.Entity("duration","val")
							hass.Entity("token","val")
)					tts.Matcher("Your [timer:timer_name] timer has been set, starting now."), 
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
				micSensor.Entities(
					"time","val"
				)
)			r.Match( 
				riemann.Group( 
					timers.Matcher("timers", 
							hass.Entity("duration","val")
							hass.Entity("name","val")
							hass.Entity("time","val")
					timeout.Matcher("timeout", 
							hass.Entity("duration","val")
							hass.Entity("token","val")
)					tts.Matcher("Your [timer:timer_name] timer has been set, starting now."), 
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
				micSensor.Entities(
					"timer_name","val"
					"duration","val"
				)
)			r.Match( 
				riemann.Group( 
					timers.Matcher("timers", 
							hass.Entity("duration","val")
							hass.Entity("name","val")
							hass.Entity("time","val")
					timeout.Matcher("timeout", 
							hass.Entity("duration","val")
							hass.Entity("token","val")
)					tts.Matcher("Your [timer:timer_name] timer has been set, starting now."), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("affirm"),
				micSensor.Intents("affirm:affirm"),
				micSensor.Entities(
					"timer_name","val"
					"duration","val"
				)
)			r.Match( 
				riemann.Group( 
					mic.Matcher("mic", 
					timers.Matcher("timers", 
							hass.Entity("duration","val")
							hass.Entity("name","val")
							hass.Entity("time","val")
					timeout.Matcher("timeout", 
							hass.Entity("duration","val")
							hass.Entity("token","val")
)					tts.Matcher("You have a timer called [timer:timer_name]. Are you sure you want to make another?Your [timer:timer_name] timer has been set, starting now."), 
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
				micSensor.Entities(
					"timer_name","val"
					"duration","val"
				)
)			r.Match( 
				riemann.Group( 
					mic.Matcher("mic", 
					timers.Matcher("timers", 
							hass.Entity("name","val")
)					tts.Matcher("You have a timer called [timer:timer_name]. Are you sure you want to make another?Okay, I won't create it."), 
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
				micSensor.Entities(
					"timer_name","val"
					"time","val"
				)
)			r.Match( 
				riemann.Group( 
					timers.Matcher("timers", 
							hass.Entity("duration","val")
							hass.Entity("name","val")
							hass.Entity("time","val")
					timeout.Matcher("timeout", 
							hass.Entity("duration","val")
							hass.Entity("token","val")
)					tts.Matcher("Your [timer:timer_name] timer has been set, starting now."), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("affirm"),
				micSensor.Intents("affirm:affirm"),
				micSensor.Entities(
					"timer_name","val"
					"time","val"
				)
)			r.Match( 
				riemann.Group( 
					mic.Matcher("mic", 
					timers.Matcher("timers", 
							hass.Entity("duration","val")
							hass.Entity("name","val")
							hass.Entity("time","val")
					timeout.Matcher("timeout", 
							hass.Entity("duration","val")
							hass.Entity("token","val")
)					tts.Matcher("You have a timer called [timer:timer_name]. Are you sure you want to make another?Your [timer:timer_name] timer has been set, starting now."), 
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
				micSensor.Entities(
					"timer_name","val"
					"time","val"
				)
)			r.Match( 
				riemann.Group( 
					mic.Matcher("mic", 
					timers.Matcher("timers", 
							hass.Entity("name","val")
)					tts.Matcher("You have a timer called [timer:timer_name]. Are you sure you want to make another?Okay, I won't create it."), 
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
)			r.Match( 
				riemann.Group( 
					timers.Matcher("timers", 
							hass.Entity("duration","val")
							hass.Entity("name","val")
							hass.Entity("time","val")
)					tts.Matcher("You have no [timer:timer_name] timers set."), 
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
)			r.Match( 
				riemann.Group( 
					timers.Matcher("timers", 
)				), 
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
				micSensor.Entities(
					"timer_name","val"
				)
)			r.Match( 
				riemann.Group( 
					timers.Matcher("timers", 
							hass.Entity("duration","val")
							hass.Entity("name","val")
							hass.Entity("time","val")
)					tts.Matcher("You have no [timer:timer_name] timers set."), 
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
				micSensor.Entities(
					"timer_name","val"
				)
)			r.Match( 
				riemann.Group( 
					timers.Matcher("timers", 
)				), 
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
				micSensor.Entities(
					"duration","val"
				)
)			r.Match( 
				riemann.Group( 
					timers.Matcher("timers", 
							hass.Entity("duration","val")
							hass.Entity("name","val")
							hass.Entity("time","val")
)					tts.Matcher("You have no [timer:timer_name] timers set."), 
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
				micSensor.Entities(
					"duration","val"
				)
)			r.Match( 
				riemann.Group( 
					timers.Matcher("timers", 
)				), 
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
				micSensor.Entities(
					"time","val"
				)
)			r.Match( 
				riemann.Group( 
					timers.Matcher("timers", 
							hass.Entity("duration","val")
							hass.Entity("name","val")
							hass.Entity("time","val")
)					tts.Matcher("You have no [timer:timer_name] timers set."), 
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
				micSensor.Entities(
					"time","val"
				)
)			r.Match( 
				riemann.Group( 
					timers.Matcher("timers", 
)				), 
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
				micSensor.Entities(
					"timer_name","val"
					"duration","val"
				)
)			r.Match( 
				riemann.Group( 
					timers.Matcher("timers", 
							hass.Entity("duration","val")
							hass.Entity("name","val")
							hass.Entity("time","val")
)					tts.Matcher("You have no [timer:timer_name] timers set."), 
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
				micSensor.Entities(
					"timer_name","val"
					"duration","val"
				)
)			r.Match( 
				riemann.Group( 
					timers.Matcher("timers", 
)				), 
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
				micSensor.Entities(
					"method","val"
					"time","val"
					"timer_name","val"
				)
)			r.Match( 
				riemann.Group( 
					timers.Matcher("timers", 
							hass.Entity("duration","val")
							hass.Entity("name","val")
							hass.Entity("time","val")
)					tts.Matcher("You have no [timer:timer_name] timers set."), 
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
				micSensor.Entities(
					"method","val"
					"time","val"
					"timer_name","val"
				)
)			r.Match( 
				riemann.Group( 
					timers.Matcher("timers", 
)				), 
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
)			r.Match( 
				riemann.Group( 
					timers.Matcher("timers", 
							hass.Entity("name","val")
)					tts.Matcher("You have no [timer:timer_name] timers set."), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("affirm"),
				micSensor.Intents("affirm:affirm"),
)			r.Match( 
				riemann.Group( 
					mic.Matcher("mic", 
					timers.Matcher("timers", 
)					tts.Matcher("Are you sure you want to delete your [timer:timer_name] timer?Deleting your [timer:timer_name] timer."), 
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
)			r.Match( 
				riemann.Group( 
					mic.Matcher("mic", 
					timers.Matcher("timers", 
							hass.Entity("name","val")
)					tts.Matcher("Are you sure you want to delete your [timer:timer_name] timer?Okay, I won't delete it."), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("affirm"),
				micSensor.Intents("affirm:affirm"),
)			r.Match( 
				riemann.Group( 
					mic.Matcher("mic", 
					timers.Matcher("timers", 
)					tts.Matcher("Are you sure you want to delete your [timer:timer_name] timers?Deleting your [timer:timer_name] timers."), 
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
)			r.Match( 
				riemann.Group( 
					mic.Matcher("mic", 
					timers.Matcher("timers", 
							hass.Entity("name","val")
)					tts.Matcher("Are you sure you want to delete your [timer:timer_name] timers?Okay, I won't delete them."), 
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
				micSensor.Entities(
					"timer_name","val"
				)
)			r.Match( 
				riemann.Group( 
					timers.Matcher("timers", 
							hass.Entity("name","val")
)					tts.Matcher("You have no [timer:timer_name] timers set."), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("affirm"),
				micSensor.Intents("affirm:affirm"),
				micSensor.Entities(
					"timer_name","val"
				)
)			r.Match( 
				riemann.Group( 
					mic.Matcher("mic", 
					timers.Matcher("timers", 
)					tts.Matcher("Are you sure you want to delete your [timer:timer_name] timer?Deleting your [timer:timer_name] timer."), 
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
				micSensor.Entities(
					"timer_name","val"
				)
)			r.Match( 
				riemann.Group( 
					mic.Matcher("mic", 
					timers.Matcher("timers", 
							hass.Entity("name","val")
)					tts.Matcher("Are you sure you want to delete your [timer:timer_name] timer?Okay, I won't delete it."), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("affirm"),
				micSensor.Intents("affirm:affirm"),
				micSensor.Entities(
					"timer_name","val"
				)
)			r.Match( 
				riemann.Group( 
					mic.Matcher("mic", 
					timers.Matcher("timers", 
)					tts.Matcher("Are you sure you want to delete your [timer:timer_name] timers?Deleting your [timer:timer_name] timers."), 
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
				micSensor.Entities(
					"timer_name","val"
				)
)			r.Match( 
				riemann.Group( 
					mic.Matcher("mic", 
					timers.Matcher("timers", 
							hass.Entity("name","val")
)					tts.Matcher("Are you sure you want to delete your [timer:timer_name] timers?Okay, I won't delete them."), 
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
				micSensor.Entities(
					"duration","val"
				)
)			r.Match( 
				riemann.Group( 
					timers.Matcher("timers", 
							hass.Entity("name","val")
)					tts.Matcher("You have no [timer:timer_name] timers set."), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("affirm"),
				micSensor.Intents("affirm:affirm"),
				micSensor.Entities(
					"duration","val"
				)
)			r.Match( 
				riemann.Group( 
					mic.Matcher("mic", 
					timers.Matcher("timers", 
)					tts.Matcher("Are you sure you want to delete your [timer:timer_name] timer?Deleting your [timer:timer_name] timer."), 
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
				micSensor.Entities(
					"duration","val"
				)
)			r.Match( 
				riemann.Group( 
					mic.Matcher("mic", 
					timers.Matcher("timers", 
							hass.Entity("name","val")
)					tts.Matcher("Are you sure you want to delete your [timer:timer_name] timer?Okay, I won't delete it."), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("affirm"),
				micSensor.Intents("affirm:affirm"),
				micSensor.Entities(
					"duration","val"
				)
)			r.Match( 
				riemann.Group( 
					mic.Matcher("mic", 
					timers.Matcher("timers", 
)					tts.Matcher("Are you sure you want to delete your [timer:timer_name] timers?Deleting your [timer:timer_name] timers."), 
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
				micSensor.Entities(
					"duration","val"
				)
)			r.Match( 
				riemann.Group( 
					mic.Matcher("mic", 
					timers.Matcher("timers", 
							hass.Entity("name","val")
)					tts.Matcher("Are you sure you want to delete your [timer:timer_name] timers?Okay, I won't delete them."), 
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
				micSensor.Entities(
					"time","val"
				)
)			r.Match( 
				riemann.Group( 
					timers.Matcher("timers", 
							hass.Entity("name","val")
)					tts.Matcher("You have no [timer:timer_name] timers set."), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("affirm"),
				micSensor.Intents("affirm:affirm"),
				micSensor.Entities(
					"time","val"
				)
)			r.Match( 
				riemann.Group( 
					mic.Matcher("mic", 
					timers.Matcher("timers", 
)					tts.Matcher("Are you sure you want to delete your [timer:timer_name] timer?Deleting your [timer:timer_name] timer."), 
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
				micSensor.Entities(
					"time","val"
				)
)			r.Match( 
				riemann.Group( 
					mic.Matcher("mic", 
					timers.Matcher("timers", 
							hass.Entity("name","val")
)					tts.Matcher("Are you sure you want to delete your [timer:timer_name] timer?Okay, I won't delete it."), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("affirm"),
				micSensor.Intents("affirm:affirm"),
				micSensor.Entities(
					"time","val"
				)
)			r.Match( 
				riemann.Group( 
					mic.Matcher("mic", 
					timers.Matcher("timers", 
)					tts.Matcher("Are you sure you want to delete your [timer:timer_name] timers?Deleting your [timer:timer_name] timers."), 
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
				micSensor.Entities(
					"time","val"
				)
)			r.Match( 
				riemann.Group( 
					mic.Matcher("mic", 
					timers.Matcher("timers", 
							hass.Entity("name","val")
)					tts.Matcher("Are you sure you want to delete your [timer:timer_name] timers?Okay, I won't delete them."), 
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
				micSensor.Entities(
					"timer_name","val"
					"duration","val"
				)
)			r.Match( 
				riemann.Group( 
					timers.Matcher("timers", 
							hass.Entity("name","val")
)					tts.Matcher("You have no [timer:timer_name] timers set."), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("affirm"),
				micSensor.Intents("affirm:affirm"),
				micSensor.Entities(
					"timer_name","val"
					"duration","val"
				)
)			r.Match( 
				riemann.Group( 
					mic.Matcher("mic", 
					timers.Matcher("timers", 
)					tts.Matcher("Are you sure you want to delete your [timer:timer_name] timer?Deleting your [timer:timer_name] timer."), 
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
				micSensor.Entities(
					"timer_name","val"
					"duration","val"
				)
)			r.Match( 
				riemann.Group( 
					mic.Matcher("mic", 
					timers.Matcher("timers", 
							hass.Entity("name","val")
)					tts.Matcher("Are you sure you want to delete your [timer:timer_name] timer?Okay, I won't delete it."), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("affirm"),
				micSensor.Intents("affirm:affirm"),
				micSensor.Entities(
					"timer_name","val"
					"duration","val"
				)
)			r.Match( 
				riemann.Group( 
					mic.Matcher("mic", 
					timers.Matcher("timers", 
)					tts.Matcher("Are you sure you want to delete your [timer:timer_name] timers?Deleting your [timer:timer_name] timers."), 
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
				micSensor.Entities(
					"timer_name","val"
					"duration","val"
				)
)			r.Match( 
				riemann.Group( 
					mic.Matcher("mic", 
					timers.Matcher("timers", 
							hass.Entity("name","val")
)					tts.Matcher("Are you sure you want to delete your [timer:timer_name] timers?Okay, I won't delete them."), 
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
				micSensor.Entities(
					"timer_name","val"
					"time","val"
				)
)			r.Match( 
				riemann.Group( 
					timers.Matcher("timers", 
							hass.Entity("name","val")
)					tts.Matcher("You have no [timer:timer_name] timers set."), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("affirm"),
				micSensor.Intents("affirm:affirm"),
				micSensor.Entities(
					"timer_name","val"
					"time","val"
				)
)			r.Match( 
				riemann.Group( 
					mic.Matcher("mic", 
					timers.Matcher("timers", 
)					tts.Matcher("Are you sure you want to delete your [timer:timer_name] timer?Deleting your [timer:timer_name] timer."), 
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
				micSensor.Entities(
					"timer_name","val"
					"time","val"
				)
)			r.Match( 
				riemann.Group( 
					mic.Matcher("mic", 
					timers.Matcher("timers", 
							hass.Entity("name","val")
)					tts.Matcher("Are you sure you want to delete your [timer:timer_name] timer?Okay, I won't delete it."), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("affirm"),
				micSensor.Intents("affirm:affirm"),
				micSensor.Entities(
					"timer_name","val"
					"time","val"
				)
)			r.Match( 
				riemann.Group( 
					mic.Matcher("mic", 
					timers.Matcher("timers", 
)					tts.Matcher("Are you sure you want to delete your [timer:timer_name] timers?Deleting your [timer:timer_name] timers."), 
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
				micSensor.Entities(
					"timer_name","val"
					"time","val"
				)
)			r.Match( 
				riemann.Group( 
					mic.Matcher("mic", 
					timers.Matcher("timers", 
							hass.Entity("name","val")
)					tts.Matcher("Are you sure you want to delete your [timer:timer_name] timers?Okay, I won't delete them."), 
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
)			r.Match( 
				riemann.Group( 
					mic.Matcher("mic", 
							hass.Entity("entity","val")
					timers.Matcher("timers", 
							hass.Entity("duration","val")
							hass.Entity("name","val")
							hass.Entity("time","val")
					timeout.Matcher("timeout", 
							hass.Entity("duration","val")
							hass.Entity("token","val")
)					tts.Matcher("Okay, for how long?Your [timer:timer_name] timer has been set, starting now."), 
				), 
			) 
		}) 
	}) 
