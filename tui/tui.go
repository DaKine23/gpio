package tui

import (
	"github.com/DaKine23/gpio/gpio"
	tb "github.com/nsf/termbox-go"
)

var background = tb.ColorBlack
var textColor = tb.ColorWhite

type TUI struct {
}

//Draw is only a delegate to termbox-go.Flush()
func (o *TUI) Draw() error {
	err := tb.Flush()
	if err != nil {
		return err
	}
	return nil
}

var maxX, maxY int

func (o *TUI) DrawLedStrip(ledStrip *gpio.GPIO_LedSet, offset, columns int, isHorizontal bool, color, color2 tb.Attribute) {
	x, y := tb.Size()
	stripSize := len(ledStrip.Set)

	colLength := stripSize / columns
	if stripSize%2 != 0 {
		colLength++
	}
	if isHorizontal {

		distance := x / stripSize * columns
		for k, v := range ledStrip.Set {
			o.DrawLed(distance/2+(k%colLength)*distance, offset*((k/colLength)+1), color, color2, bool(v.Value), v.Selected)
		}
	} else {

		distance := y / stripSize * columns
		for k, v := range ledStrip.Set {
			o.DrawLed(offset*((k/colLength)+1), distance/2+(k%colLength)*distance, color, color2, bool(v.Value), v.Selected)
		}
	}

}

func (o *TUI) DrawLed(x, y int, color, color2 tb.Attribute, isOn, isSelected bool) {

	maxX, maxY = tb.Size()

	if x > maxX {
		x = maxX - 1
	}
	if x < 0 {
		x = 0 + 1
	}
	if y > maxY {
		y = maxY - 1
	}
	if y < 0 {
		y = 0 + 1
	}
	tb.SetCell(x-1, y, ' ', tb.ColorWhite, color)
	tb.SetCell(x, y-1, ' ', tb.ColorWhite, color)
	tb.SetCell(x+1, y, ' ', tb.ColorWhite, color)
	tb.SetCell(x, y+1, ' ', tb.ColorWhite, color)
	if isOn {

		tb.SetCell(x, y, ' ', tb.ColorWhite, tb.ColorWhite)
	} else {

		tb.SetCell(x, y, ' ', tb.ColorWhite, color2)
	}
	if isSelected {

		tb.SetCell(x+1, y+1, ' ', tb.ColorWhite, tb.ColorYellow)
		tb.SetCell(x-1, y+1, ' ', tb.ColorWhite, tb.ColorYellow)
		tb.SetCell(x-1, y-1, ' ', tb.ColorWhite, tb.ColorYellow)
		tb.SetCell(x+1, y-1, ' ', tb.ColorWhite, tb.ColorYellow)
	} else {
		tb.SetCell(x+1, y+1, ' ', tb.ColorDefault, tb.ColorDefault)
		tb.SetCell(x-1, y+1, ' ', tb.ColorDefault, tb.ColorDefault)
		tb.SetCell(x-1, y-1, ' ', tb.ColorDefault, tb.ColorDefault)
		tb.SetCell(x+1, y-1, ' ', tb.ColorDefault, tb.ColorDefault)

	}

}
