package handlers

import (
	"github.com/gin-gonic/gin"
	"nashor/internal/helpers"
	"nashor/internal/services"
)

func HandleGetSummonerById(c *gin.Context) {
	var (
		puuid  = c.Param("puuid")
		region = c.Param("region")
	)

	data, err := services.GetSummonerByPuuid(region, puuid)

	if err != nil {
		res := helpers.HttpResFromErr(err)

		c.JSON(res.Status, res)
		return
	}

	c.JSON(200, data)
}
