package handlers

import (
	"github.com/gin-gonic/gin"
	"nashor/internal/problem"
	"nashor/internal/services"
	"os"
)

func HandleGetAccountByRiotID(c *gin.Context) {
	var (
		region   = os.Getenv("REGION")
		gameName = c.Param("gameName")
		tagLine  = c.Param("tagLine")
	)

	data, err := services.GetAccountByRiotId(region, gameName, tagLine)

	if err != nil {
		var res problem.ErrorResponse
		res.FromErr(err)

		c.JSON(res.Status, res)
		return
	}

	c.JSON(200, data)
}
