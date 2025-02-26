package handlers

import (
	"github.com/gin-gonic/gin"
	"nashor/internal/helpers"
	"nashor/internal/services"
	"os"
)

func HandleGetAccountByRiotID(c *gin.Context) {
	var (
		gameName = c.Param("gameName")
		tagLine  = c.Param("tagLine")
	)

	data, err := services.GetAccountByRiotId(os.Getenv("REGION"), gameName, tagLine)

	if err != nil {
		res := helpers.HttpResFromErr(err)

		c.JSON(res.Status, res)
		return
	}

	c.JSON(200, data)
}
