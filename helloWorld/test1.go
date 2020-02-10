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

			drone.TakeOff()
			fmt.Println("Listo")
			drone.Forward(100)
			drone.Clockwise(100)

			gobot.Every(2*time.Second, func() {
				drone.Forward(100)
				drone.Right(100)
				drone.Clockwise(100)
			})

		}
	
	robot := gobot.NewRobot("tello",
		[]gobot.Connection{},
		[]gobot.Device{drone},
		work,
	)
	robot.Start()
}