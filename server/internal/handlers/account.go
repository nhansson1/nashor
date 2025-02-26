package handlers

import (
	"github.com/gin-gonic/gin"
	"nashor/internal/helpers"
	"nashor/internal/services"
)

func HandleGetAccountByRiotID(c *gin.Context) {
	gameName := c.Param("gameName")
	tagLine := c.Param("tagLine")

	data, err := services.GetAccountByRiotId(gameName, tagLine)

	if err != nil {
		res := helpers.HttpResFromErr(err)

		c.JSON(res.Status, res)
		return
	}

	c.JSON(200, data)
}
