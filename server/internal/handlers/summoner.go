package handlers

import (
	"github.com/gin-gonic/gin"
	"nashor/internal/helpers"
	"nashor/internal/services"
)

func HandleGetSummonerById(c *gin.Context) {
	puuid := c.Param("puuid")

	data, err := services.GetSummonerByPuuid(puuid)

	if err != nil {
		res := helpers.HttpResFromErr(err)

		c.JSON(res.Status, res)
		return
	}

	c.JSON(200, data)
}
