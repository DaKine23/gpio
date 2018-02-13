package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DaKine23/gpio/data"
	"github.com/DaKine23/gpio/tui"
	"github.com/gin-gonic/gin"
	tb "github.com/nsf/termbox-go"
)

var tu *tui.TUI

func test() {
	fmt.Println("vim-go")
}

func All(c *gin.Context) {

	c.JSON(http.StatusOK, data.Strip)

}

func Add(c *gin.Context) {
	port := c.Param("port")

	data.Strip.Add(port)
	data.Strip.Write("")

	tu.DrawLedStrip(data.Strip, data.Offset, data.Colnum, true, tb.ColorBlue, tb.ColorBlack)
	tu.Draw()
	bs, _ := json.Marshal(data.Strip)
	if bs != nil {
		data.M.Broadcast(bs)
	}

}

func Remove(c *gin.Context) {
	data.Strip.RemoveSelected()
	data.Strip.Write("")
	tu.DrawLedStrip(data.Strip, data.Offset, data.Colnum, true, tb.ColorBlue, tb.ColorBlack)
	tu.Draw()
	bs, _ := json.Marshal(data.Strip)
	if bs != nil {
		data.M.Broadcast(bs)
	}
}

func Next(c *gin.Context) {
	data.Strip.SelectNext(1, true)
	data.Strip.Write("")
	tu.DrawLedStrip(data.Strip, data.Offset, data.Colnum, true, tb.ColorBlue, tb.ColorBlack)
	tu.Draw()
	bs, _ := json.Marshal(data.Strip)
	if bs != nil {
		data.M.Broadcast(bs)
	}
}

func Previous(c *gin.Context) {
	data.Strip.SelectNext(1, false)
	data.Strip.Write("")
	tu.DrawLedStrip(data.Strip, data.Offset, data.Colnum, true, tb.ColorBlue, tb.ColorBlack)
	tu.Draw()
	bs, _ := json.Marshal(data.Strip)
	if bs != nil {
		data.M.Broadcast(bs)
	}
}

func MoveRight(c *gin.Context) {

	data.Strip.Move(true)
	data.Strip.Write("")
	tu.DrawLedStrip(data.Strip, data.Offset, data.Colnum, true, tb.ColorBlue, tb.ColorBlack)
	tu.Draw()
	bs, _ := json.Marshal(data.Strip)
	if bs != nil {
		data.M.Broadcast(bs)
	}

}
func MoveLeft(c *gin.Context) {
	data.Strip.Move(false)
	data.Strip.Write("")
	tu.DrawLedStrip(data.Strip, data.Offset, data.Colnum, true, tb.ColorBlue, tb.ColorBlack)
	tu.Draw()
	bs, _ := json.Marshal(data.Strip)
	if bs != nil {
		data.M.Broadcast(bs)
	}

}

func SwitchSelected(c *gin.Context) {
	data.Strip.SwitchSelected()
	data.Strip.Write("")
	tu.DrawLedStrip(data.Strip, data.Offset, data.Colnum, true, tb.ColorBlue, tb.ColorBlack)
	tu.Draw()
	bs, _ := json.Marshal(data.Strip)
	if bs != nil {
		data.M.Broadcast(bs)
	}

}

func SwitchAll(c *gin.Context) {
	data.Strip.AllSwitch()
	data.Strip.Write("")
	tu.DrawLedStrip(data.Strip, data.Offset, data.Colnum, true, tb.ColorBlue, tb.ColorBlack)
	tu.Draw()
	bs, _ := json.Marshal(data.Strip)
	if bs != nil {
		data.M.Broadcast(bs)
	}
}
