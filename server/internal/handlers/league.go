package handlers

import (
	"github.com/gin-gonic/gin"
	"nashor/internal/helpers"
	"nashor/internal/services"
)

func HandleGetLeagueEntriesById(c *gin.Context) {
	summonerId := c.Param("summonerId")

	data, err := services.GetRankQueusById(summonerId)

	if err != nil {
		res := helpers.HttpResFromErr(err)

		c.JSON(res.Status, res)
		return
	}

	c.JSON(200, data)
}
