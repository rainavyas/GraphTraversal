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
				micSensor.Domains("weather"),
				micSensor.Intents("weather:query_condition"),
			)
			r.Match( 
				riemann.Group( 
					tts.Matcher("The weather is [weather:description] and the temperature is [weather:temperature_value] degrees"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number1", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("weather"),
				micSensor.Intents("weather:query_condition"),
				micSensor.Entities(
					"location","Tokyo"
				)
			)
			r.Match( 
				riemann.Group( 
					tts.Matcher("In [weather:location] the weather is [weather:description] and the temperature is [weather:temperature_value] degrees"), 
				), 
			) 
		}) 
	}) 
	Describe("Test number2", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("weather"),
				micSensor.Intents("weather:query_condition"),
				micSensor.Entities(
					"time_range","today"
				)
			)
			r.Match( 
				riemann.Group( 
					tts.Matcher("In [weather:location] the weather [weather:weekday] is [weather:description] with highs of [weather:max_temperature_value] and lows of [weather:min_temperature_value]."), 
				), 
			) 
		}) 
	}) 
	Describe("Test number3", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("weather"),
				micSensor.Intents("weather:query_condition"),
				micSensor.Entities(
					"time_range","today"
					"location","Tokyo"
				)
			)
			r.Match( 
				riemann.Group( 
					tts.Matcher("In [weather:location] the weather [weather:weekday] is [weather:description] with highs of [weather:max_temperature_value] and lows of [weather:min_temperature_value]."), 
				), 
			) 
		}) 
	}) 
	Describe("Test number4", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("weather"),
				micSensor.Intents("weather:query_condition"),
				micSensor.Entities(
					"time_range","tomorrow"
				)
			)
			r.Match( 
				riemann.Group( 
					tts.Matcher("In [weather:location] the weather [weather:weekday] is [weather:description] with highs of [weather:max_temperature_value] and lows of [weather:min_temperature_value]."), 
				), 
			) 
		}) 
	}) 
	Describe("Test number5", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("weather"),
				micSensor.Intents("weather:query_condition"),
				micSensor.Entities(
					"time_range","tomorrow"
					"location","Tokyo"
				)
			)
			r.Match( 
				riemann.Group( 
					tts.Matcher("In [weather:location] the weather [weather:weekday] is [weather:description] with highs of [weather:max_temperature_value] and lows of [weather:min_temperature_value]."), 
				), 
			) 
		}) 
	}) 
	Describe("Test number6", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("weather"),
				micSensor.Intents("weather:query_condition"),
				micSensor.Entities(
					"time_range","this week"
				)
			)
			r.Match( 
				riemann.Group( 
					tts.Matcher("Here's the weather in [weather:location]!"), 
					tts.Matcher("[weather:weekday] it will be [weather:description] with highs of [weather:max_temperature_value] and lows of [weather:min_temperature_value]."), 
					tts.Matcher("[weather:weekday] it will be [weather:description] with highs of [weather:max_temperature_value] and lows of [weather:min_temperature_value]."), 
					tts.Matcher("[weather:weekday] it will be [weather:description] with highs of [weather:max_temperature_value] and lows of [weather:min_temperature_value]."), 
					tts.Matcher("[weather:weekday] it will be [weather:description] with highs of [weather:max_temperature_value] and lows of [weather:min_temperature_value]."), 
				), 
			) 
		}) 
	}) 
	Describe("Test number7", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("weather"),
				micSensor.Intents("weather:query_condition"),
				micSensor.Entities(
					"time_range","this week"
				)
			)
			r.Match( 
				riemann.Group( 
					tts.Matcher("Here's the weather in [weather:location]!"), 
					tts.Matcher("[weather:weekday] it will be [weather:description] with highs of [weather:max_temperature_value] and lows of [weather:min_temperature_value]."), 
					tts.Matcher("[weather:weekday] it will be [weather:description] with highs of [weather:max_temperature_value] and lows of [weather:min_temperature_value]."), 
					tts.Matcher("[weather:weekday] it will be [weather:description] with highs of [weather:max_temperature_value] and lows of [weather:min_temperature_value]."), 
				), 
			) 
		}) 
	}) 
	Describe("Test number8", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("weather"),
				micSensor.Intents("weather:query_condition"),
				micSensor.Entities(
					"time_range","this week"
				)
			)
			r.Match( 
				riemann.Group( 
					tts.Matcher("Here's the weather in [weather:location]!"), 
					tts.Matcher("[weather:weekday] it will be [weather:description] with highs of [weather:max_temperature_value] and lows of [weather:min_temperature_value]."), 
					tts.Matcher("[weather:weekday] it will be [weather:description] with highs of [weather:max_temperature_value] and lows of [weather:min_temperature_value]."), 
				), 
			) 
		}) 
	}) 
	Describe("Test number9", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("weather"),
				micSensor.Intents("weather:query_condition"),
				micSensor.Entities(
					"time_range","this week"
				)
			)
			r.Match( 
				riemann.Group( 
					tts.Matcher("Here's the weather in [weather:location]!"), 
					tts.Matcher("[weather:weekday] it will be [weather:description] with highs of [weather:max_temperature_value] and lows of [weather:min_temperature_value]."), 
				), 
			) 
		}) 
	}) 
	Describe("Test number10", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("weather"),
				micSensor.Intents("weather:query_condition"),
				micSensor.Entities(
					"time_range","this week"
				)
			)
			r.Match( 
				riemann.Group( 
				), 
			) 
		}) 
	}) 
	Describe("Test number11", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("weather"),
				micSensor.Intents("weather:query_condition"),
				micSensor.Entities(
					"time_range","this week"
					"location","Tokyo"
				)
			)
			r.Match( 
				riemann.Group( 
					tts.Matcher("Here's the weather in [weather:location]!"), 
					tts.Matcher("[weather:weekday] it will be [weather:description] with highs of [weather:max_temperature_value] and lows of [weather:min_temperature_value]."), 
					tts.Matcher("[weather:weekday] it will be [weather:description] with highs of [weather:max_temperature_value] and lows of [weather:min_temperature_value]."), 
					tts.Matcher("[weather:weekday] it will be [weather:description] with highs of [weather:max_temperature_value] and lows of [weather:min_temperature_value]."), 
					tts.Matcher("[weather:weekday] it will be [weather:description] with highs of [weather:max_temperature_value] and lows of [weather:min_temperature_value]."), 
				), 
			) 
		}) 
	}) 
	Describe("Test number12", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("weather"),
				micSensor.Intents("weather:query_condition"),
				micSensor.Entities(
					"time_range","this week"
					"location","Tokyo"
				)
			)
			r.Match( 
				riemann.Group( 
					tts.Matcher("Here's the weather in [weather:location]!"), 
					tts.Matcher("[weather:weekday] it will be [weather:description] with highs of [weather:max_temperature_value] and lows of [weather:min_temperature_value]."), 
					tts.Matcher("[weather:weekday] it will be [weather:description] with highs of [weather:max_temperature_value] and lows of [weather:min_temperature_value]."), 
					tts.Matcher("[weather:weekday] it will be [weather:description] with highs of [weather:max_temperature_value] and lows of [weather:min_temperature_value]."), 
				), 
			) 
		}) 
	}) 
	Describe("Test number13", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("weather"),
				micSensor.Intents("weather:query_condition"),
				micSensor.Entities(
					"time_range","this week"
					"location","Tokyo"
				)
			)
			r.Match( 
				riemann.Group( 
					tts.Matcher("Here's the weather in [weather:location]!"), 
					tts.Matcher("[weather:weekday] it will be [weather:description] with highs of [weather:max_temperature_value] and lows of [weather:min_temperature_value]."), 
					tts.Matcher("[weather:weekday] it will be [weather:description] with highs of [weather:max_temperature_value] and lows of [weather:min_temperature_value]."), 
				), 
			) 
		}) 
	}) 
	Describe("Test number14", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("weather"),
				micSensor.Intents("weather:query_condition"),
				micSensor.Entities(
					"time_range","this week"
					"location","Tokyo"
				)
			)
			r.Match( 
				riemann.Group( 
					tts.Matcher("Here's the weather in [weather:location]!"), 
					tts.Matcher("[weather:weekday] it will be [weather:description] with highs of [weather:max_temperature_value] and lows of [weather:min_temperature_value]."), 
				), 
			) 
		}) 
	}) 
	Describe("Test number15", func() { 
		It("Should say Hi", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("weather"),
				micSensor.Intents("weather:query_condition"),
				micSensor.Entities(
					"time_range","this week"
					"location","Tokyo"
				)
			)
			r.Match( 
				riemann.Group( 
				), 
			) 
		}) 
	}) 
