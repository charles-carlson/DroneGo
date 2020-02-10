package main

import (
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dji/tello"
)

var (
	drone      = tello.NewDriver("8080")
	flightData *tello.FlightData
)

func main() {

	work := func() {
		drone.On(tello.FlightDataEvent, func(data interface{}) {
			flightData = data.(*tello.FlightData)
		})

		drone.TakeOff()
		drone.SetVector(.5, 0, .1, 0)
		gobot.Every(3*time.Second, func() {
			drone.SetPsi(.8)
		})
		gobot.Every(5*time.Second, func() {
			drone.CeaseRotation()
		})
		gobot.After(25*time.Second, func() {
			drone.Land()
		})

	}

	robot := gobot.NewRobot("tello",
		[]gobot.Connection{},
		[]gobot.Device{drone},
		work,
	)

	robot.Start()
}
