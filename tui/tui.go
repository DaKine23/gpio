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

func drawTextField(fg, bg tb.Attribute, lines ...string) {

	msgLength := make([]int, len(lines))

	tb.Clear(tb.ColorBlack, tb.ColorBlack)
	tb.Flush()
	maxMsgLength := 0
	for i := range lines {
		msgLength[i] = len(lines[i])
		if msgLength[i] > maxMsgLength {
			maxMsgLength = msgLength[i]
		}
	}

	// draw 6 lines of color

	x, y := tb.Size()
	i := (x - maxMsgLength + 2) / 2
	j := (y - len(lines) + 2) / 2
	counteri, counterj := 0, 0

	for counteri < maxMsgLength+2 {

		for counterj < len(lines)+2 {

			if counterj > 0 && counterj <= len(lines) && counteri > 0 && counteri <= msgLength[counterj-1] {
				tb.SetCell(i, j, rune(lines[counterj-1][counteri-1]), fg, bg)
			} else {
				tb.SetCell(i, j, ' ', fg, bg)
			}

			counterj++
			j++
		}
		j = (y - len(lines) + 2) / 2

		counterj = 0

		counteri++
		i++
	}
	tb.Flush()

}

func (o *TUI) TextRequestModal(question string, fg, bg tb.Attribute) string {

	answer := ""
	var event tb.Event
	for event.Key != tb.KeyEsc && event.Key != tb.KeyEnter {
		drawTextField(fg, bg, question, "", answer)
		event = tb.PollEvent()

		if event.Key == tb.KeyBackspace || event.Key == tb.KeyBackspace2 {
			if len(answer) > 0 {
				answer = answer[:len(answer)-1]
			}
		} else if event.Key != tb.KeySpace {
			answer += string(event.Ch)
		}

	}

	if event.Key == tb.KeyEnter {
		return answer
	}

	return ""
}

func (o *TUI) DrawLedStrip(ledStrip *gpio.GPIO_LedSet, offset, columns int, isHorizontal bool, color, color2 tb.Attribute) {
	tb.Clear(tb.ColorBlack, tb.ColorBlack)
	tb.Flush()

	if len(ledStrip.Set) == 0 {
		return
	}

	x, y := tb.Size()
	stripSize := len(ledStrip.Set)
	mod := 0
	colLength := stripSize / columns
	if stripSize%2 != 0 && columns%2 != 1 {
		colLength++
		mod = 1
	}
	if isHorizontal {

		distance := x / (stripSize + mod) * columns
		for k, v := range ledStrip.Set {
			o.DrawLed(distance/2+(k%colLength)*distance, offset*((k/colLength)+1), color, color2, bool(v.Value), v.Selected, v.Port)
		}
	} else {

		distance := y / stripSize * columns
		for k, v := range ledStrip.Set {
			o.DrawLed(offset*((k/colLength)+1), distance/2+(k%colLength)*distance, color, color2, bool(v.Value), v.Selected, v.Port)
		}
	}

}

func (o *TUI) DrawLed(x, y int, color, color2 tb.Attribute, isOn, isSelected bool, port string) {

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
	for k := 0; k < len(port); k++ {
		tb.SetCell(x-1+k, y-2, rune(port[k]), tb.ColorWhite, tb.ColorBlack)
	}

}
