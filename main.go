package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/DaKine23/gpio/gpio"
)

func main() {

	pins2 := []string{"1", "3", "2", "4", "5", "6", "7", "8"}

	delay := 0

	if len(os.Args) > 3 {

		delay, _ = strconv.Atoi(os.Args[1])

		pins2 = os.Args[2:]
	}

	i := 0
	var gp *gpio.GPIO

	strip := gp.NewLedSet(pins2...)
	strip.SingleOn("strip", "1")
	strip.SingleOn("strip", "2")
	fmt.Println(strip.Oostring())
	fmt.Println("----move r-----")

	for i < 9 {
		time.Sleep(time.Duration(delay) * time.Millisecond)
		strip.Move("move right", true)
		fmt.Println(strip.Oostring())
		i++
	}
	i = 0
	fmt.Println("----move l-----")
	for i < 9 {
		time.Sleep(time.Duration(delay) * time.Millisecond)
		strip.Move("move left", false)
		fmt.Println(strip.Oostring())
		i++
	}
	fmt.Println("----switch-----")
	time.Sleep(time.Duration(delay) * time.Millisecond)
	strip.AllSwitch("strip")
	fmt.Println(strip.Oostring())
	time.Sleep(time.Duration(delay) * time.Millisecond)
	strip.AllSwitch("strip")
	fmt.Println(strip.Oostring())

}
