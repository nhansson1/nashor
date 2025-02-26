package handlers

import (
	"github.com/gin-gonic/gin"
	"nashor/internal/helpers"
	"nashor/internal/services"
)

func HandleGetLeagueEntriesById(c *gin.Context) {
	var (
		region     = c.Param("region")
		summonerId = c.Param("summonerId")
	)

	data, err := services.GetRankQueusById(region, summonerId)

	if err != nil {
		res := helpers.HttpResFromErr(err)

		c.JSON(res.Status, res)
		return
	}

	c.JSON(200, data)
}
