package router

import (
	"io/ioutil"
	"net/http"

	"github.com/DaKine23/gpio/data"
	"github.com/DaKine23/gpio/handler"
	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
)

func Init() {

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())

	r.Static("/static", "./public/static")
	r.Static("/fonts", "./public/fonts")
	r.Static("/img", "./public/img")
	r.Static("/js", "./public/js")

	r.GET("/all", handler.All)
	r.GET("/ws", DefaultHandler)

	r.POST("/add/:port", handler.Add)
	r.POST("/remove", handler.Remove)
	r.POST("/next", handler.Next)
	r.POST("/previous", handler.Previous)
	r.POST("/move/right", handler.MoveRight)
	r.POST("/move/left", handler.MoveLeft)
	r.POST("/switch/selected", handler.SwitchSelected)
	r.POST("/switch/all", handler.SwitchAll)

	r.GET("/", func(c *gin.Context) {
		content, _ := ioutil.ReadFile("./public/index.html")

		c.Data(http.StatusOK, "text/html", content)

	})

	r.Run(":8080")

}

func SetDefaultWebsocketHandler(newMelody *melody.Melody) {
	data.M = newMelody
}

func DefaultHandler(c *gin.Context) {
	data.M.HandleRequest(c.Writer, c.Request)
}
