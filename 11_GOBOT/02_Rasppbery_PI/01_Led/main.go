package main

import (
	"time"

	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {

	r := raspi.NewAdaptor()
	gpioPin := gpio.NewDirectPinDriver(r, "3")

	for {
		gpioPin.DigitalWrite(1)
		time.Sleep(3 * time.Second)
		gpioPin.DigitalWrite(0)
		time.Sleep(3 * time.Second)
	}

	/*
		r := raspi.NewAdaptor()
		led := gpio.NewLedDriver(r, "3")

		work := func() {
			gobot.Every(1*time.Second, func() {
				led.Toggle()
			})
		}

		robot := gobot.NewRobot("blinkBot",
			[]gobot.Connection{r},
			[]gobot.Device{led},
			work,
		)

		robot.Start()*/
}

// GOARM=7 GOARCH=arm GOOS=linux go build main.go
