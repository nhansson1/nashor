package handler

import (
    "nashor/internal/problem"
    "github.com/gin-gonic/gin"
)

func (h Handler) HandleGetSummonerById(c *gin.Context) {
	var (
		puuid  = c.Param("puuid")
		region = c.Param("region")
	)

	data, err := h.summonerService.GetSummonerByPuuid(region, puuid)

	if err != nil {
		var res problem.ErrorResponse
		res.FromErr(err)

		c.JSON(res.Status, res)
		return
	}

	c.JSON(200, data)
}
