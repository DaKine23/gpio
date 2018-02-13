package main

import (
	"os"
	"strconv"

	"github.com/DaKine23/gpio/data"
	"github.com/DaKine23/gpio/router"
	"github.com/DaKine23/gpio/tui"
	tb "github.com/nsf/termbox-go"
)

var tu *tui.TUI

func main() {

	err := tb.Init()
	if err != nil {
		panic(err)
	}
	defer tb.Close()

	colnum := 1

	if len(os.Args) > 1 {
		coln, err := strconv.Atoi(os.Args[1])
		if err != nil {
			panic(os.Args[1])
		}
		colnum = coln
	}
	data.Init()
	//data.Strip.Set[0].Selected = true

	go router.Init()

	tu.DrawLedStrip(data.Strip, 4, colnum, true, tb.ColorBlue, tb.ColorBlack)
	tu.Draw()
	var event tb.Event
	odd := 0
	if len(data.Strip.Set)%2 != 0 {
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

			data.Strip.Add(port)
			tu.DrawLedStrip(data.Strip, 4, colnum, true, tb.ColorBlue, tb.ColorBlack)
			tu.Draw()
			data.Strip.Write("")

		}
		if event.Ch == '-' {

			data.Strip.RemoveSelected()
			tu.DrawLedStrip(data.Strip, 4, colnum, true, tb.ColorBlue, tb.ColorBlack)
			tu.Draw()
			data.Strip.Write("")

		}

		// (CR)U(D)

		if event.Key == tb.KeyArrowRight {
			data.Strip.SelectNext(1, true)
			tu.DrawLedStrip(data.Strip, 4, colnum, true, tb.ColorBlue, tb.ColorBlack)
			tu.Draw()
			data.Strip.Write("")

		}
		if event.Key == tb.KeyArrowLeft {
			data.Strip.SelectNext(1, false)
			tu.DrawLedStrip(data.Strip, 4, colnum, true, tb.ColorBlue, tb.ColorBlack)
			tu.Draw()
			data.Strip.Write("")
		}
		if event.Key == tb.KeyArrowDown {

			data.Strip.SelectNext(len(data.Strip.Set)/colnum+odd, true)
			tu.DrawLedStrip(data.Strip, 4, colnum, true, tb.ColorBlue, tb.ColorBlack)
			tu.Draw()
			data.Strip.Write("")
		}
		if event.Key == tb.KeyArrowUp {
			data.Strip.SelectNext(len(data.Strip.Set)/colnum+odd, false)
			tu.DrawLedStrip(data.Strip, 4, colnum, true, tb.ColorBlue, tb.ColorBlack)
			tu.Draw()
			data.Strip.Write("")
		}
		if event.Key == tb.KeyEnter {
			data.Strip.SwitchSelected()
			tu.DrawLedStrip(data.Strip, 4, colnum, true, tb.ColorBlue, tb.ColorBlack)
			tu.Draw()
			data.Strip.Write("")
		}
		if event.Key == tb.KeyPgup {
			data.Strip.Move(true)
			tu.DrawLedStrip(data.Strip, 4, colnum, true, tb.ColorBlue, tb.ColorBlack)
			tu.Draw()
			data.Strip.Write("")
		}
		if event.Key == tb.KeyPgdn {
			data.Strip.Move(false)
			tu.DrawLedStrip(data.Strip, 4, colnum, true, tb.ColorBlue, tb.ColorBlack)
			tu.Draw()
			data.Strip.Write("")
		}
		if event.Key == tb.KeyEnd {
			data.Strip.AllSwitch()
			tu.DrawLedStrip(data.Strip, 4, colnum, true, tb.ColorBlue, tb.ColorBlack)
			tu.Draw()
			data.Strip.Write("")
		}
	}

}
