package data

import (
	"os"

	"github.com/DaKine23/gpio/gpio"
	"github.com/olahol/melody"
)

var gp *gpio.GPIO

var Strip *gpio.GPIO_LedSet
var Colnum int = 1
var Offset int = 4

var M *melody.Melody

func Init() {

	pins2 := []string{}

	if len(os.Args) > 2 {

		pins2 = os.Args[2:]

	}

	Strip = gp.NewLedSet(pins2...)
	Strip.Add("10")

}
