package main

//
//import (
//	"time"
//
//	"github.com/hybridgroup/gobot"
//	"github.com/hybridgroup/gobot/platforms/gpio"
//	"github.com/hybridgroup/gobot/platforms/raspi"
//)
//
//func main() {
//	gbot := gobot.NewGobot()
//
//	r := raspi.NewRaspiAdaptor("raspi")
//	led := gpio.NewLedDriver(r, "led", "17")
//
//	work := func() {
//		gobot.Every(1*time.Second, func() {
//			led.Toggle()
//		})
//	}
//
//	robot := gobot.NewRobot("blinkBot",
//		[]gobot.Connection{r},
//		[]gobot.Device{led},
//		work,
//	)
//
//	gbot.AddRobot(robot)
//
//	gbot.Start()
//}
//

import (
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"os"
	"time"
)

func main() {

	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer rpio.Close()

	pin := rpio.Pin(17)
	pin.Output()

	for x := 0; x < 20; x++ {
		pin.Toggle()
		time.Sleep(time.Second)
	}
}