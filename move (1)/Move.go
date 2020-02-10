package main

import (
	"fmt"
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
		gobot.After(10*time.Second, func() {
			drone.TakeOff()
			drone.Hover()
			fmt.Println("Ready")
		})

		gobot.Every(6*time.Second, func() {
			drone.Forward(100)
			gobot.Every(7*time.Second, func() {
				drone.Clockwise(45)
				gobot.After(20*time.Millisecond, func() {
					drone.CounterClockwise(45)
				})
			})
		})

		gobot.After(20*time.Second, func() {
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
