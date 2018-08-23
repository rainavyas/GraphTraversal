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

	Describe("Test number0", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_on"),
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_on", 
							hass.Entity("color","green")
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
					)
					tts.Matcher("There is no light called like that"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number1", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_on"),
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_on", 
							hass.Entity("color","green")
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
					)
					tts.Matcher("turning the lights on"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number2", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_on"),
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_on", 
							hass.Entity("color","green")
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
					)
					lights.Matcher("turn_on", 
							hass.Entity("color","green")
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
					)
					tts.Matcher("There is no light called like that"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number3", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_on"),
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_on", 
							hass.Entity("color","green")
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
					)
					lights.Matcher("turn_on", 
							hass.Entity("color","green")
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
					)
					tts.Matcher("turning the lights on"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number4", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_on"),
				micSensor.Entities(
					"house_place","bedroom"
				)
			)
			r.Match( 
				riemann.Group( 
					tts.Matcher("There is no light called like that"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number5", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_on"),
				micSensor.Entities(
					"house_place","bedroom"
				)
			)
			r.Match( 
				riemann.Group( 
					tts.Matcher("turning the lights on"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number6", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_on"),
				micSensor.Entities(
					"house_place","bedroom"
				)
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_on", 
							hass.Entity("color","green")
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
					)
					tts.Matcher("There is no light called like that"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number7", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_on"),
				micSensor.Entities(
					"house_place","bedroom"
				)
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_on", 
							hass.Entity("color","green")
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
					)
					tts.Matcher("turning the lights on"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number8", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_on"),
				micSensor.Entities(
					"house_place","bedroom"
				)
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_on", 
							hass.Entity("color","green")
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
					)
					lights.Matcher("turn_on", 
							hass.Entity("color","green")
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
					)
					tts.Matcher("There is no light called like that"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number9", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_on"),
				micSensor.Entities(
					"house_place","bedroom"
				)
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_on", 
							hass.Entity("color","green")
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
					)
					lights.Matcher("turn_on", 
							hass.Entity("color","green")
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
					)
					tts.Matcher("turning the lights on"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number10", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_on"),
				micSensor.Entities(
					"house_place","bedroom"
					"color","green"
				)
			)
			r.Match( 
				riemann.Group( 
					tts.Matcher("There is no light called like that"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number11", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_on"),
				micSensor.Entities(
					"house_place","bedroom"
					"color","green"
				)
			)
			r.Match( 
				riemann.Group( 
					tts.Matcher("turning the lights on"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number12", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_on"),
				micSensor.Entities(
					"house_place","bedroom"
					"color","green"
				)
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_on", 
							hass.Entity("color","green")
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
					)
					tts.Matcher("There is no light called like that"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number13", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_on"),
				micSensor.Entities(
					"house_place","bedroom"
					"color","green"
				)
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_on", 
							hass.Entity("color","green")
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
					)
					tts.Matcher("turning the lights on"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number14", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_on"),
				micSensor.Entities(
					"house_place","bedroom"
					"color","green"
				)
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_on", 
							hass.Entity("color","green")
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
					)
					lights.Matcher("turn_on", 
							hass.Entity("color","green")
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
					)
					tts.Matcher("There is no light called like that"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number15", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_on"),
				micSensor.Entities(
					"house_place","bedroom"
					"color","green"
				)
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_on", 
							hass.Entity("color","green")
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
					)
					lights.Matcher("turn_on", 
							hass.Entity("color","green")
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
					)
					tts.Matcher("turning the lights on"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number16", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_up"),
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_on", 
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
							hass.Entity("brightness","255")
					)
					tts.Matcher("There is no light called like that"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number17", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_up"),
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_on", 
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
							hass.Entity("brightness","255")
					)
					tts.Matcher("brightening the lights"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number18", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_up"),
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_on", 
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
							hass.Entity("brightness","255")
					)
					lights.Matcher("turn_on", 
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
							hass.Entity("brightness","255")
					)
					tts.Matcher("There is no light called like that"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number19", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_up"),
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_on", 
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
							hass.Entity("brightness","255")
					)
					lights.Matcher("turn_on", 
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
							hass.Entity("brightness","255")
					)
					tts.Matcher("brightening the lights"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number20", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_up"),
				micSensor.Entities(
					"house_place","bedroom"
				)
			)
			r.Match( 
				riemann.Group( 
					tts.Matcher("There is no light called like that"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number21", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_up"),
				micSensor.Entities(
					"house_place","bedroom"
				)
			)
			r.Match( 
				riemann.Group( 
					tts.Matcher("brightening the lights"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number22", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_up"),
				micSensor.Entities(
					"house_place","bedroom"
				)
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_on", 
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
							hass.Entity("brightness","255")
					)
					tts.Matcher("There is no light called like that"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number23", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_up"),
				micSensor.Entities(
					"house_place","bedroom"
				)
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_on", 
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
							hass.Entity("brightness","255")
					)
					tts.Matcher("brightening the lights"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number24", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_up"),
				micSensor.Entities(
					"house_place","bedroom"
				)
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_on", 
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
							hass.Entity("brightness","255")
					)
					lights.Matcher("turn_on", 
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
							hass.Entity("brightness","255")
					)
					tts.Matcher("There is no light called like that"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number25", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_up"),
				micSensor.Entities(
					"house_place","bedroom"
				)
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_on", 
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
							hass.Entity("brightness","255")
					)
					lights.Matcher("turn_on", 
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
							hass.Entity("brightness","255")
					)
					tts.Matcher("brightening the lights"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number26", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_down"),
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_on", 
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
							hass.Entity("brightness","255")
					)
					tts.Matcher("There is no light called like that"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number27", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_down"),
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_on", 
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
							hass.Entity("brightness","255")
					)
					tts.Matcher("dimming the lights"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number28", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_down"),
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_on", 
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
							hass.Entity("brightness","255")
					)
					lights.Matcher("turn_on", 
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
							hass.Entity("brightness","255")
					)
					tts.Matcher("There is no light called like that"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number29", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_down"),
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_on", 
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
							hass.Entity("brightness","255")
					)
					lights.Matcher("turn_on", 
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
							hass.Entity("brightness","255")
					)
					tts.Matcher("dimming the lights"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number30", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_down"),
				micSensor.Entities(
					"house_place","bedroom"
				)
			)
			r.Match( 
				riemann.Group( 
					tts.Matcher("There is no light called like that"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number31", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_down"),
				micSensor.Entities(
					"house_place","bedroom"
				)
			)
			r.Match( 
				riemann.Group( 
					tts.Matcher("dimming the lights"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number32", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_down"),
				micSensor.Entities(
					"house_place","bedroom"
				)
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_on", 
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
							hass.Entity("brightness","255")
					)
					tts.Matcher("There is no light called like that"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number33", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_down"),
				micSensor.Entities(
					"house_place","bedroom"
				)
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_on", 
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
							hass.Entity("brightness","255")
					)
					tts.Matcher("dimming the lights"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number34", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_down"),
				micSensor.Entities(
					"house_place","bedroom"
				)
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_on", 
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
							hass.Entity("brightness","255")
					)
					lights.Matcher("turn_on", 
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
							hass.Entity("brightness","255")
					)
					tts.Matcher("There is no light called like that"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number35", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_down"),
				micSensor.Entities(
					"house_place","bedroom"
				)
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_on", 
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
							hass.Entity("brightness","255")
					)
					lights.Matcher("turn_on", 
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
							hass.Entity("brightness","255")
					)
					tts.Matcher("dimming the lights"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number36", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:set_specific"),
				micSensor.Entities(
					"house_place","bedroom"
					"color","green"
				)
			)
			r.Match( 
				riemann.Group( 
					tts.Matcher("There is no light called like that"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number37", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:set_specific"),
				micSensor.Entities(
					"house_place","bedroom"
					"color","green"
				)
			)
			r.Match( 
				riemann.Group( 
					tts.Matcher("changing the colour of the lights"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number38", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:set_specific"),
				micSensor.Entities(
					"house_place","bedroom"
					"color","green"
				)
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_on", 
							hass.Entity("color","green")
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
					)
					tts.Matcher("There is no light called like that"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number39", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:set_specific"),
				micSensor.Entities(
					"house_place","bedroom"
					"color","green"
				)
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_on", 
							hass.Entity("color","green")
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
					)
					tts.Matcher("changing the colour of the lights"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number40", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:set_specific"),
				micSensor.Entities(
					"house_place","bedroom"
					"color","green"
				)
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_on", 
							hass.Entity("color","green")
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
					)
					lights.Matcher("turn_on", 
							hass.Entity("color","green")
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
					)
					tts.Matcher("There is no light called like that"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number41", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:set_specific"),
				micSensor.Entities(
					"house_place","bedroom"
					"color","green"
				)
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_on", 
							hass.Entity("color","green")
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
					)
					lights.Matcher("turn_on", 
							hass.Entity("color","green")
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
					)
					tts.Matcher("changing the colour of the lights"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number42", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:set_specific"),
				micSensor.Entities(
					"color","green"
				)
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_on", 
							hass.Entity("color","green")
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
					)
					tts.Matcher("There is no light called like that"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number43", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:set_specific"),
				micSensor.Entities(
					"color","green"
				)
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_on", 
							hass.Entity("color","green")
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
					)
					tts.Matcher("changing the colour of the lights"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number44", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:set_specific"),
				micSensor.Entities(
					"color","green"
				)
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_on", 
							hass.Entity("color","green")
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
					)
					lights.Matcher("turn_on", 
							hass.Entity("color","green")
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
					)
					tts.Matcher("There is no light called like that"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number45", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:set_specific"),
				micSensor.Entities(
					"color","green"
				)
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_on", 
							hass.Entity("color","green")
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
					)
					lights.Matcher("turn_on", 
							hass.Entity("color","green")
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
					)
					tts.Matcher("changing the colour of the lights"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number46", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_off"),
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_off", 
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
					)
					tts.Matcher("There is no light called like that"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number47", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_off"),
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_off", 
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
					)
					tts.Matcher("turning the lights off"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number48", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_off"),
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_off", 
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
					)
					lights.Matcher("turn_off", 
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
					)
					tts.Matcher("There is no light called like that"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number49", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_off"),
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_off", 
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
					)
					lights.Matcher("turn_off", 
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
					)
					tts.Matcher("turning the lights off"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number50", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_off"),
				micSensor.Entities(
					"house_place","bedroom"
				)
			)
			r.Match( 
				riemann.Group( 
					tts.Matcher("There is no light called like that"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number51", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_off"),
				micSensor.Entities(
					"house_place","bedroom"
				)
			)
			r.Match( 
				riemann.Group( 
					tts.Matcher("turning the lights off"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number52", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_off"),
				micSensor.Entities(
					"house_place","bedroom"
				)
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_off", 
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
					)
					tts.Matcher("There is no light called like that"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number53", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_off"),
				micSensor.Entities(
					"house_place","bedroom"
				)
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_off", 
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
					)
					tts.Matcher("turning the lights off"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number54", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_off"),
				micSensor.Entities(
					"house_place","bedroom"
				)
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_off", 
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
					)
					lights.Matcher("turn_off", 
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
					)
					tts.Matcher("There is no light called like that"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number55", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:turn_off"),
				micSensor.Entities(
					"house_place","bedroom"
				)
			)
			r.Match( 
				riemann.Group( 
					lights.Matcher("turn_off", 
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
					)
					lights.Matcher("turn_off", 
							hass.Entity("entity_id","light.gateway_light_34ce008be244")
					)
					tts.Matcher("turning the lights off"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number56", func() { 
		It("Should say Hi", func() { 
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
	Describe("Test number57", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("iot_lighting"),
				micSensor.Intents("iot_lighting:set_specific"),
				micSensor.Entities(
					"house_place","bedroom"
				)
			)
			r.Match( 
				riemann.Group( 
					tts.Matcher("Could not find the color you ask for."), 
				), 
			) 
		}) 
	}) 
