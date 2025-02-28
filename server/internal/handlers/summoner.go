package handlers

import (
	"github.com/gin-gonic/gin"
	"nashor/internal/problem"
	"nashor/internal/services"
)

func HandleGetSummonerById(c *gin.Context) {
	var (
		puuid  = c.Param("puuid")
		region = c.Param("region")
	)

	data, err := services.GetSummonerByPuuid(region, puuid)

	if err != nil {
		var res problem.ErrorResponse
		res.FromErr(err)

		c.JSON(res.Status, res)
		return
	}

	c.JSON(200, data)
}
