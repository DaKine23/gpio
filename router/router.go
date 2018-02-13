package router

import (
	"io/ioutil"
	"net/http"

	"github.com/DaKine23/gpio/handler"
	"github.com/gin-gonic/gin"
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

	r.GET("/", func(c *gin.Context) {
		content, _ := ioutil.ReadFile("./public/index.html")

		c.Data(http.StatusOK, "text/html", content)

	})

	r.Run(":8080")

}
