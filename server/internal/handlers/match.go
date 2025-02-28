package handlers

import (
	"github.com/gin-gonic/gin"
	"nashor/internal/problem"
	"nashor/internal/services"
	"os"
)

func HandleGetMatchesByPuuid(c *gin.Context) {
	var (
		region = os.Getenv("REGION")
		puuid  = c.Param("puuid")
		start  = c.Param("start")
		count  = c.Param("count")
	)

	ids, err := services.GetMatchIdsByPuuid(region, puuid, start, count)

	if err != nil {
		var perr problem.ErrorResponse
		perr.FromErr(err)

		c.JSON(perr.Status, perr)
		return
	}

	matches, err := services.GetMatchDataByIds(region, ids)

	if err != nil {
		var perr problem.ErrorResponse
		perr.FromErr(err)

		c.JSON(perr.Status, perr)
		return
	}

	c.JSON(200, matches)
}
