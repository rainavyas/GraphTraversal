package tests

import (
	"flag"
	"net/http"
	"time"

	"bitbucket.org/emotech/common/golang/protos"
	"bitbucket.org/emotech/riemann/actuators/camera"
	"bitbucket.org/emotech/riemann/actuators/leds"
	micActuator "bitbucket.org/emotech/riemann/actuators/mic"
	"bitbucket.org/emotech/riemann/actuators/motors"
	"bitbucket.org/emotech/riemann/actuators/tts"
	"bitbucket.org/emotech/riemann/dolly"
	"bitbucket.org/emotech/riemann/iot/hass"
	"bitbucket.org/emotech/riemann/iot/lights"
	"bitbucket.org/emotech/riemann/loglines"
	"bitbucket.org/emotech/riemann/riemann"
	micSensor "bitbucket.org/emotech/riemann/sensors/mic"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

var dollyAddress = flag.String("dolly", "http://dolly:8080", "dolly address")

var _ = Describe("Riemann", func() {
	gomega.SetDefaultEventuallyTimeout(10. * time.Second)

	var olly *dolly.Dolly
	var r *riemann.Riemann
	BeforeSuite(func() {
		olly = dolly.NewDolly(*dollyAddress, &http.Client{})
		r = riemann.NewRiemann()
		r.Run()
		logger := loglines.NewLoggerWrapper(GinkgoWriter)
		logger.Start()
	})

	BeforeEach(func() {
		r.Restart(olly)
		r.WaitFor(olly).ToBe(riemann.StartedAndIdle())
	})

	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("recognize_user"),
				micSensor.Intents("recognize_user:recognize_user"),
				micSensor.Entities(
					"current_user_name","val"
				)
)			r.Match( 
				riemann.Group( 
					mic.Matcher("mic", 
)					tts.Matcher("how are you doinghello, [current_user_name]"), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("recognize_user"),
				micSensor.Intents("recognize_user:recognize_user"),
)			r.Match( 
				riemann.Group( 
					mic.Matcher("mic", 
)					tts.Matcher("how are you doingI don't think we've been introduced"), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("calendar"),
				micSensor.Intents("calendar:calendar"),
				micSensor.Entities(
					"date","val"
				)
)			r.Match( 
				riemann.Group( 
					mic.Matcher("mic", 
)					tts.Matcher("how are you doingI don't know you enough to know your schedule"), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("calendar"),
				micSensor.Intents("calendar:calendar"),
				micSensor.Entities(
					"date","val"
				)
)			r.Match( 
				riemann.Group( 
					mic.Matcher("mic", 
)					tts.Matcher("how are you doingI don't know you enough to know your schedule"), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("calendar"),
				micSensor.Intents("calendar:calendar"),
				micSensor.Entities(
					"current_user_id","val"
					"date","val"
				)
)			r.Match( 
				riemann.Group( 
					mic.Matcher("mic", 
					calendar.Matcher("calendar", 
							hass.Entity("from_time","val")
)					tts.Matcher("how are you doinglet's have a look [current_user_name]I couldn't find any events"), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("calendar"),
				micSensor.Intents("calendar:calendar"),
				micSensor.Entities(
					"current_user_id","val"
					"date","val"
				)
)			r.Match( 
				riemann.Group( 
					mic.Matcher("mic", 
					calendar.Matcher("calendar", 
)					tts.Matcher("how are you doinglet's have a look [current_user_name]You only have one event for [date]."), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("calendar"),
				micSensor.Intents("calendar:calendar"),
				micSensor.Entities(
					"current_user_id","val"
					"date","val"
				)
)			r.Match( 
				riemann.Group( 
					mic.Matcher("mic", 
					calendar.Matcher("calendar", 
)					tts.Matcher("how are you doinglet's have a look [current_user_name]You are a bit busy [date], you have [length] events. "), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("calendar"),
				micSensor.Intents("calendar:calendar"),
				micSensor.Entities(
					"current_user_id","val"
					"date","val"
				)
)			r.Match( 
				riemann.Group( 
					mic.Matcher("mic", 
					calendar.Matcher("calendar", 
)					tts.Matcher("how are you doinglet's have a look [current_user_name]You are a bit busy [date], you have [length] events. "), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("calendar"),
				micSensor.Intents("calendar:calendar"),
				micSensor.Entities(
					"current_user_id","val"
					"date","val"
				)
)			r.Match( 
				riemann.Group( 
					mic.Matcher("mic", 
					calendar.Matcher("calendar", 
							hass.Entity("from_time","val")
)					tts.Matcher("how are you doinglet's have a look [current_user_name]You look really busy! you have [length] events for [date]"), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("calendar"),
				micSensor.Intents("calendar:calendar"),
				micSensor.Entities(
					"current_user_id","val"
					"date","val"
				)
)			r.Match( 
				riemann.Group( 
					mic.Matcher("mic", 
					calendar.Matcher("calendar", 
							hass.Entity("from_time","val")
)					tts.Matcher("how are you doinglet's have a look [current_user_name]I couldn't find any events"), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("calendar"),
				micSensor.Intents("calendar:calendar"),
				micSensor.Entities(
					"current_user_id","val"
					"date","val"
				)
)			r.Match( 
				riemann.Group( 
					mic.Matcher("mic", 
					calendar.Matcher("calendar", 
)					tts.Matcher("how are you doinglet's have a look [current_user_name]You only have one event for [date]."), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("calendar"),
				micSensor.Intents("calendar:calendar"),
				micSensor.Entities(
					"current_user_id","val"
					"date","val"
				)
)			r.Match( 
				riemann.Group( 
					mic.Matcher("mic", 
					calendar.Matcher("calendar", 
)					tts.Matcher("how are you doinglet's have a look [current_user_name]You are a bit busy [date], you have [length] events. "), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("calendar"),
				micSensor.Intents("calendar:calendar"),
				micSensor.Entities(
					"current_user_id","val"
					"date","val"
				)
)			r.Match( 
				riemann.Group( 
					mic.Matcher("mic", 
					calendar.Matcher("calendar", 
)					tts.Matcher("how are you doinglet's have a look [current_user_name]You are a bit busy [date], you have [length] events. "), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("calendar"),
				micSensor.Intents("calendar:calendar"),
				micSensor.Entities(
					"current_user_id","val"
					"date","val"
				)
)			r.Match( 
				riemann.Group( 
					mic.Matcher("mic", 
					calendar.Matcher("calendar", 
							hass.Entity("from_time","val")
)					tts.Matcher("how are you doinglet's have a look [current_user_name]You look really busy! you have [length] events for [date]"), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("+"),
				micSensor.Intents("+:+"),
)			r.Match( 
				riemann.Group( 
					mic.Matcher("mic", 
)					tts.Matcher("how are you doingSorry I didn't understand that."), 
				), 
			) 
		}) 
	}) 
