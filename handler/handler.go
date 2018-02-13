package handler

import (
	"fmt"
	"net/http"

	"github.com/DaKine23/gpio/data"
	"github.com/gin-gonic/gin"
)

func test() {
	fmt.Println("vim-go")
}

func All(c *gin.Context) {

	c.JSON(http.StatusOK, data.Strip)

}
