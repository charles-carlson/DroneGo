package main

import (
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dji/tello"
)

var (
	drone = tello.NewDriver("8080")
	flightData *tello.FlightData
)

func main() {

	work := func() {
		drone.On(tello.FlightDataEvent, func(data interface{}) {
			flightData = data.(*tello.FlightData)
		})

		drone.TakeOff()
	
		drone.SetVectors(0.5, 0, 0.3, 0.3)

		gobot.After(10*time.Second, func() {
			drone.SetVectors(-0.5, 0, -0.3, -0.3)

		drone.Forward()
		drone.Halt()
		drone.Clockwise(90)
		drone.Forward(69)
		drone.Land()

		gobot.After(10*time.Second, func() {
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
