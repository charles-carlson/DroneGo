package main

import (
	"fmt"
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
		fmt.Println("Take Off")
		drone.Left(100)
		fmt.Println("Left")
		
		gobot.After(3*time.Second, func() {
			drone.Left(0)
			drone.Forward(100)
			fmt.Println("Forward")
		})

		gobot.After(7*time.Second, func() {
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
