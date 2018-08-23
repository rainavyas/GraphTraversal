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

	Describe("User says a", func() { 
		It("Should say b", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_up"),
				micSensor.Entities(
					"house_place","val"
				)
			)
			r.Match( 
				riemann.Group( 
					iot.Matcher("resolve_by_domain_and_name", 
							hass.Entity("","val")
							hass.Entity("","val")
					)
					iot.Matcher("resolve_by_domain_name_and_capability", 
							hass.Entity("","val")
							hass.Entity("","val")
							hass.Entity("","val")
					)
					tts.Matcher("brightening the lights"), 
				), 
			) 
		}) 
	}) 
	Describe("User says d", func() { 
		It("Should say b", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_up"),
				micSensor.Entities(
					"house_place","val"
				)
			)
			r.Match( 
				riemann.Group( 
					iot.Matcher("resolve_by_domain_and_name", 
							hass.Entity("","val")
							hass.Entity("","val")
					)
					iot.Matcher("resolve_by_domain_name_and_capability", 
							hass.Entity("","val")
							hass.Entity("","val")
							hass.Entity("","val")
					)
					tts.Matcher("brightening the lights"), 
				), 
			) 
		}) 
	}) 
	Describe("User says d", func() { 
		It("Should say b", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_up"),
				micSensor.Entities(
					"house_place","val"
				)
			)
			r.Match( 
				riemann.Group( 
					iot.Matcher("resolve_by_domain_and_name", 
							hass.Entity("","val")
							hass.Entity("","val")
					)
					tts.Matcher("There is no light called like that"), 
				), 
			) 
		}) 
	}) 
	Describe("User says b", func() { 
		It("Should say b", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_down"),
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_on", 
							hass.Entity("","val")
							hass.Entity("","val")
					)
					timeout.Matcher("set", 
							hass.Entity("","val")
							hass.Entity("","val")
					)
					iot.Matcher("resolve_by_domain_name_and_capability", 
							hass.Entity("","val")
							hass.Entity("","val")
							hass.Entity("","val")
					)
				), 
			) 
		}) 
	}) 
	Describe("User says b", func() { 
		It("Should say b", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_down"),
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_on", 
							hass.Entity("","val")
							hass.Entity("","val")
					)
					timeout.Matcher("set", 
							hass.Entity("","val")
							hass.Entity("","val")
					)
					iot.Matcher("resolve_by_domain_name_and_capability", 
							hass.Entity("","val")
							hass.Entity("","val")
							hass.Entity("","val")
					)
					tts.Matcher("$error"), 
				), 
			) 
		}) 
	}) 
	Describe("User says b", func() { 
		It("Should say b", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_down"),
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_on", 
							hass.Entity("","val")
							hass.Entity("","val")
					)
					timeout.Matcher("set", 
							hass.Entity("","val")
							hass.Entity("","val")
					)
					iot.Matcher("resolve_by_domain_name_and_capability", 
							hass.Entity("","val")
							hass.Entity("","val")
							hass.Entity("","val")
					)
					tts.Matcher("dimming the lights"), 
				), 
			) 
		}) 
	}) 
	Describe("User says b", func() { 
		It("Should say b", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_off"),
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_off", 
							hass.Entity("","val")
					)
					timeout.Matcher("set", 
							hass.Entity("","val")
							hass.Entity("","val")
					)
					tts.Matcher("turning the lights off"), 
				), 
			) 
		}) 
	}) 
	Describe("User says b", func() { 
		It("Should say b", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_off"),
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_off", 
							hass.Entity("","val")
					)
					timeout.Matcher("set", 
							hass.Entity("","val")
							hass.Entity("","val")
					)
					tts.Matcher("turning the lights off"), 
				), 
			) 
		}) 
	}) 
	Describe("User says b", func() { 
		It("Should say b", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_on"),
				micSensor.Entities(
					"color","val"
					"house_place","val"
				)
			)
			r.Match( 
				riemann.Group( 
					iot.Matcher("resolve_by_domain_and_name", 
							hass.Entity("","val")
							hass.Entity("","val")
					)
					tts.Matcher("There is no light called like that"), 
				), 
			) 
		}) 
	}) 
	Describe("User says b", func() { 
		It("Should say b", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_on"),
				micSensor.Entities(
					"house_place","val"
					"color","val"
				)
			)
			r.Match( 
				riemann.Group( 
					iot.Matcher("resolve_by_domain_and_name", 
							hass.Entity("","val")
							hass.Entity("","val")
					)
					tts.Matcher("turning the lights on"), 
				), 
			) 
		}) 
	}) 
	Describe("User says b", func() { 
		It("Should say b", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:set_specific"),
				micSensor.Entities(
					"house_place","val"
				)
			)
			r.Match( 
				riemann.Group( 
					tts.Matcher("Could not find the color you ask for."), 
				), 
			) 
		}) 
	}) 
	Describe("User says b", func() { 
		It("Should say b", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_down"),
				micSensor.Entities(
					"house_place","val"
				)
			)
			r.Match( 
				riemann.Group( 
					iot.Matcher("resolve_by_domain_and_name", 
							hass.Entity("","val")
							hass.Entity("","val")
					)
					iot.Matcher("resolve_by_domain_name_and_capability", 
							hass.Entity("","val")
							hass.Entity("","val")
							hass.Entity("","val")
					)
					tts.Matcher("dimming the lights"), 
				), 
			) 
		}) 
	}) 
	Describe("User says b", func() { 
		It("Should say b", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_down"),
				micSensor.Entities(
					"house_place","val"
				)
			)
			r.Match( 
				riemann.Group( 
					iot.Matcher("resolve_by_domain_and_name", 
							hass.Entity("","val")
							hass.Entity("","val")
					)
					iot.Matcher("resolve_by_domain_name_and_capability", 
							hass.Entity("","val")
							hass.Entity("","val")
							hass.Entity("","val")
					)
					tts.Matcher("dimming the lights"), 
				), 
			) 
		}) 
	}) 
	Describe("User says b", func() { 
		It("Should say b", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_down"),
				micSensor.Entities(
					"house_place","val"
				)
			)
			r.Match( 
				riemann.Group( 
					iot.Matcher("resolve_by_domain_and_name", 
							hass.Entity("","val")
							hass.Entity("","val")
					)
					tts.Matcher("There is no light called like that"), 
				), 
			) 
		}) 
	}) 
	Describe("User says b", func() { 
		It("Should say f", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:set_specific"),
				micSensor.Entities(
					"house_place","val"
					"color","val"
				)
			)
			r.Match( 
				riemann.Group( 
					iot.Matcher("resolve_by_domain_and_name", 
							hass.Entity("","val")
							hass.Entity("","val")
					)
					tts.Matcher("There is no light called like that"), 
				), 
			) 
		}) 
	}) 
	Describe("User says d", func() { 
		It("Should say v", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:set_specific"),
				micSensor.Entities(
					"color","val"
					"house_place","val"
				)
			)
			r.Match( 
				riemann.Group( 
					iot.Matcher("resolve_by_domain_and_name", 
							hass.Entity("","val")
							hass.Entity("","val")
					)
					iot.Matcher("resolve_by_domain_name_and_capability", 
							hass.Entity("","val")
							hass.Entity("","val")
							hass.Entity("","val")
					)
					tts.Matcher("changing the colour of the lights"), 
				), 
			) 
		}) 
	}) 
	Describe("User says f", func() { 
		It("Should say d", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:set_specific"),
				micSensor.Entities(
					"color","val"
					"house_place","val"
				)
			)
			r.Match( 
				riemann.Group( 
					iot.Matcher("resolve_by_domain_and_name", 
							hass.Entity("","val")
							hass.Entity("","val")
					)
					iot.Matcher("resolve_by_domain_name_and_capability", 
							hass.Entity("","val")
							hass.Entity("","val")
							hass.Entity("","val")
					)
					tts.Matcher("changing the colour of the lights"), 
				), 
			) 
		}) 
	}) 
	Describe("User says d", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_off"),
				micSensor.Entities(
					"house_place","val"
				)
			)
			r.Match( 
				riemann.Group( 
					iot.Matcher("resolve_by_domain_and_name", 
							hass.Entity("","val")
							hass.Entity("","val")
					)
					tts.Matcher("There is no light called like that"), 
				), 
			) 
		}) 
	}) 
	Describe("User says d", func() { 
		It("Should say v", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_off"),
				micSensor.Entities(
					"house_place","val"
				)
			)
			r.Match( 
				riemann.Group( 
					iot.Matcher("resolve_by_domain_and_name", 
							hass.Entity("","val")
							hass.Entity("","val")
					)
					tts.Matcher("turning the lights off"), 
				), 
			) 
		}) 
	}) 
	Describe("User says f", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_on"),
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_on", 
							hass.Entity("","val")
							hass.Entity("","val")
					)
					timeout.Matcher("set", 
							hass.Entity("","val")
							hass.Entity("","val")
					)
					tts.Matcher("turning the lights on"), 
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
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_on", 
							hass.Entity("","val")
							hass.Entity("","val")
					)
					timeout.Matcher("set", 
							hass.Entity("","val")
							hass.Entity("","val")
					)
					tts.Matcher("turning the lights on"), 
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
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_on", 
							hass.Entity("","val")
							hass.Entity("","val")
					)
					timeout.Matcher("set", 
							hass.Entity("","val")
							hass.Entity("","val")
					)
					iot.Matcher("resolve_by_domain_name_and_capability", 
							hass.Entity("","val")
							hass.Entity("","val")
							hass.Entity("","val")
					)
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
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_on", 
							hass.Entity("","val")
							hass.Entity("","val")
					)
					timeout.Matcher("set", 
							hass.Entity("","val")
							hass.Entity("","val")
					)
					iot.Matcher("resolve_by_domain_name_and_capability", 
							hass.Entity("","val")
							hass.Entity("","val")
							hass.Entity("","val")
					)
					tts.Matcher("$error"), 
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
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_on", 
							hass.Entity("","val")
							hass.Entity("","val")
					)
					timeout.Matcher("set", 
							hass.Entity("","val")
							hass.Entity("","val")
					)
					iot.Matcher("resolve_by_domain_name_and_capability", 
							hass.Entity("","val")
							hass.Entity("","val")
							hass.Entity("","val")
					)
					tts.Matcher("brightening the lights"), 
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
				micSensor.Entities(
					"color","val"
				)
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_on", 
							hass.Entity("","val")
							hass.Entity("","val")
					)
					timeout.Matcher("set", 
							hass.Entity("","val")
							hass.Entity("","val")
					)
					lights.Matcher("turn_on", 
							hass.Entity("","val")
							hass.Entity("","val")
					)
					timeout.Matcher("set", 
							hass.Entity("","val")
							hass.Entity("","val")
					)
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
				micSensor.Entities(
					"color","val"
				)
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_on", 
							hass.Entity("","val")
							hass.Entity("","val")
					)
					timeout.Matcher("set", 
							hass.Entity("","val")
							hass.Entity("","val")
					)
					iot.Matcher("resolve_by_domain_name_and_capability", 
							hass.Entity("","val")
							hass.Entity("","val")
							hass.Entity("","val")
					)
					tts.Matcher("$error"), 
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
				micSensor.Entities(
					"color","val"
				)
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_on", 
							hass.Entity("","val")
							hass.Entity("","val")
					)
					timeout.Matcher("set", 
							hass.Entity("","val")
							hass.Entity("","val")
					)
					iot.Matcher("resolve_by_domain_name_and_capability", 
							hass.Entity("","val")
							hass.Entity("","val")
							hass.Entity("","val")
					)
					tts.Matcher("changing the colour of the lights"), 
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
			)
			r.Match( 
				riemann.Group( 
					tts.Matcher("Could not find the color you ask for."), 
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
				micSensor.Entities(
					"house_place","val"
				)
			)
			r.Match( 
				riemann.Group( 
					iot.Matcher("resolve_by_domain_and_name", 
							hass.Entity("","val")
							hass.Entity("","val")
					)
					tts.Matcher("turning the lights on"), 
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
				micSensor.Entities(
					"house_place","val"
				)
			)
			r.Match( 
				riemann.Group( 
					iot.Matcher("resolve_by_domain_and_name", 
							hass.Entity("","val")
							hass.Entity("","val")
					)
					tts.Matcher("There is no light called like that"), 
				), 
			) 
		}) 
	}) 
