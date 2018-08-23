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
				micSensor.Domains("time"),
				micSensor.Intents("time:query"),
			)
			r.Match( 
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
				micSensor.Entities(
					"location","Tokyo"
				)
			)
			r.Match( 
				riemann.Group( 
					tts.Matcher("it is [time] in [location]"), 
				), 
			) 
		}) 
	}) 
