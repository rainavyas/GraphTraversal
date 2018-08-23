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
				micSensor.Domains("music"),
				micSensor.Intents("music:music"),
				micSensor.Entities(
					"media_action","val"
				)
			)
			r.Match( 
				riemann.Group( 
					tts.Matcher("Here's some music!"), 
					media_player.Matcher("play", 
							hass.Entity("player_name","val")
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
				micSensor.Domains("music"),
				micSensor.Intents("music:music"),
				micSensor.Entities(
					"media_action","val"
				)
			)
			r.Match( 
				riemann.Group( 
					tts.Matcher("stopping the music"), 
					media_player.Matcher("stop", 
							hass.Entity("player_name","val")
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
				micSensor.Domains("music"),
				micSensor.Intents("music:music"),
				micSensor.Entities(
					"media_action","val"
				)
			)
			r.Match( 
				riemann.Group( 
					tts.Matcher("pause"), 
					media_player.Matcher("pause", 
							hass.Entity("player_name","val")
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
				micSensor.Domains("music"),
				micSensor.Intents("music:music"),
				micSensor.Entities(
					"media_action","val"
				)
			)
			r.Match( 
				riemann.Group( 
					tts.Matcher("next"), 
					media_player.Matcher("next_song", 
							hass.Entity("player_name","val")
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
				micSensor.Domains("music"),
				micSensor.Intents("music:music"),
				micSensor.Entities(
					"media_action","val"
				)
			)
			r.Match( 
				riemann.Group( 
					tts.Matcher("previous song"), 
					media_player.Matcher("previous_song", 
							hass.Entity("player_name","val")
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
				micSensor.Domains("music"),
				micSensor.Intents("music:music"),
				micSensor.Entities(
					"volume","val"
				)
			)
			r.Match( 
				riemann.Group( 
					tts.Matcher("turning up"), 
					media_player.Matcher("volume_up", 
							hass.Entity("player_name","val")
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
				micSensor.Domains("music"),
				micSensor.Intents("music:music"),
				micSensor.Entities(
					"volume","val"
				)
			)
			r.Match( 
				riemann.Group( 
					tts.Matcher("turning down"), 
					media_player.Matcher("volume_down", 
							hass.Entity("player_name","val")
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
				micSensor.Domains("music"),
				micSensor.Intents("music:music"),
				micSensor.Entities(
					"volume","val"
				)
			)
			r.Match( 
				riemann.Group( 
					tts.Matcher("mute"), 
					media_player.Matcher("volume_mute", 
							hass.Entity("player_name","val")
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
				micSensor.Domains("music"),
				micSensor.Intents("music:music"),
				micSensor.Entities(
					"media_action","val"
				)
			)
			r.Match( 
				riemann.Group( 
					tts.Matcher("I'm not playing anything at the moment"), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("music"),
				micSensor.Intents("music:music"),
				micSensor.Entities(
					"currently_playing","val"
					"media_action","val"
				)
			)
			r.Match( 
				riemann.Group( 
					tts.Matcher("I'm not playing anything at the moment"), 
				), 
			) 
		}) 
	}) 
	Describe("User says ", func() { 
		It("Should say ", func() { 
			r.Sensors.Mic.SendVT(micSensor.WAKEUP) 
			r.WaitFor(olly).ToBe(riemann.Listening()) 
			r.Sensors.Mic.SendNlpResult( 
				micSensor.Domains("music"),
				micSensor.Intents("music:music"),
				micSensor.Entities(
					"currently_playing","val"
					"media_action","val"
				)
			)
			r.Match( 
				riemann.Group( 
					tts.Matcher("This is [currently_playing]"), 
				), 
			) 
		}) 
	}) 
