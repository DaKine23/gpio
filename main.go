package main

import (
	"os"
	"strconv"

	"github.com/DaKine23/gpio/gpio"
	"github.com/DaKine23/gpio/tui"
	tb "github.com/nsf/termbox-go"
)

var tu *tui.TUI
var gp *gpio.GPIO

func main() {

	err := tb.Init()
	if err != nil {
		panic(err)
	}
	defer tb.Close()

	pins2 := []string{}

	colnum := 1

	if len(os.Args) > 1 {
		coln, err := strconv.Atoi(os.Args[1])
		if err != nil {
			panic(os.Args[1])
		}
		colnum = coln
	}

	if len(os.Args) > 2 {

		pins2 = os.Args[2:]

	}

	strip := gp.NewLedSet(pins2...)
	strip.Set[0].Selected = true
	tu.DrawLedStrip(strip, 4, colnum, true, tb.ColorBlue, tb.ColorBlack)
	tu.Draw()
	var event tb.Event
	odd := 0
	if len(strip.Set)%2 != 0 {
		odd++
	}
	for event.Key != tb.KeyEsc {
		event = tb.PollEvent()
		// control existing led

		// C(RU)D

		if event.Ch == '+' {

			port := tu.TextRequestModal("Please Enter the Port for the new GPIO", tb.ColorWhite, tb.ColorGreen)
			tb.Clear(tb.ColorBlack, tb.ColorBlack)
			tb.Flush()

			strip.Add(port)
			tu.DrawLedStrip(strip, 4, colnum, true, tb.ColorBlue, tb.ColorBlack)
			tu.Draw()
			strip.Write("")

		}
		if event.Ch == '-' {

			strip.RemoveSelected()
			tu.DrawLedStrip(strip, 4, colnum, true, tb.ColorBlue, tb.ColorBlack)
			tu.Draw()
			strip.Write("")

		}

		// (CR)U(D)

		if event.Key == tb.KeyArrowRight {
			strip.SelectNext(1, true)
			tu.DrawLedStrip(strip, 4, colnum, true, tb.ColorBlue, tb.ColorBlack)
			tu.Draw()
			strip.Write("")

		}
		if event.Key == tb.KeyArrowLeft {
			strip.SelectNext(1, false)
			tu.DrawLedStrip(strip, 4, colnum, true, tb.ColorBlue, tb.ColorBlack)
			tu.Draw()
			strip.Write("")
		}
		if event.Key == tb.KeyArrowDown {

			strip.SelectNext(len(strip.Set)/colnum+odd, true)
			tu.DrawLedStrip(strip, 4, colnum, true, tb.ColorBlue, tb.ColorBlack)
			tu.Draw()
			strip.Write("")
		}
		if event.Key == tb.KeyArrowUp {
			strip.SelectNext(len(strip.Set)/colnum+odd, false)
			tu.DrawLedStrip(strip, 4, colnum, true, tb.ColorBlue, tb.ColorBlack)
			tu.Draw()
			strip.Write("")
		}
		if event.Key == tb.KeyEnter {
			strip.SwitchSelected()
			tu.DrawLedStrip(strip, 4, colnum, true, tb.ColorBlue, tb.ColorBlack)
			tu.Draw()
			strip.Write("")
		}
		if event.Key == tb.KeyPgup {
			strip.Move(true)
			tu.DrawLedStrip(strip, 4, colnum, true, tb.ColorBlue, tb.ColorBlack)
			tu.Draw()
			strip.Write("")
		}
		if event.Key == tb.KeyPgdn {
			strip.Move(false)
			tu.DrawLedStrip(strip, 4, colnum, true, tb.ColorBlue, tb.ColorBlack)
			tu.Draw()
			strip.Write("")
		}
		if event.Key == tb.KeyEnd {
			strip.AllSwitch()
			tu.DrawLedStrip(strip, 4, colnum, true, tb.ColorBlue, tb.ColorBlack)
			tu.Draw()
			strip.Write("")
		}
	}

}
