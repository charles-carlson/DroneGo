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
	ry         float32
	rx         float32
	ly         float32
	lx         float32
)

func main() {

	work := func() {
		drone.On(tello.FlightDataEvent, func(data interface{}) {
			flightData = data.(*tello.FlightData)
		})

		drone.TakeOff()

		ry, rx, ly, lx = drone.Vector()
		fmt.Println("%f\n", ry, rx, ly, lx)

		drone.Forward(100)
		gobot.After(3*time.Second, func() {
			drone.SetVector(0.7, .5, 0.1, 0.8)
		})
		gobot.After(6*time.Second, func() {
			drone.CeaseRotation()
			drone.SetY(0)
			drone.Forward(100)
		})

		ry, rx, ly, lx = drone.Vector()
		fmt.Println("%f\n", ry, rx, ly, lx)

	}

	robot := gobot.NewRobot("tello",
		[]gobot.Connection{},
		[]gobot.Device{drone},
		work,
	)

	robot.Start()
}
