package data

import (
	"os"

	"github.com/DaKine23/gpio/gpio"
)

var gp *gpio.GPIO

var Strip *gpio.GPIO_LedSet

func Init() {

	pins2 := []string{}

	if len(os.Args) > 2 {

		pins2 = os.Args[2:]

	}

	Strip = gp.NewLedSet(pins2...)
	Strip.Add("10")

}
