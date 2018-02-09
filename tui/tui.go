package tui

import (
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

func (o *TUI) DrawLed(x, y int, color, color2 tb.Attribute, isOn bool) {

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

}
