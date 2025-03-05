package handler

import (
    "nashor/internal/problem"
    "github.com/gin-gonic/gin"
)

func (h Handler) HandleGetLeagueEntriesById(c *gin.Context) {
	var (
		region     = c.Param("region")
		summonerId = c.Param("summonerId")
	)

	data, err := h.leagueService.GetRankQueusById(region, summonerId)

	if err != nil {
		var res problem.ErrorResponse
		res.FromErr(err)

		c.JSON(res.Status, res)
		return
	}

	c.JSON(200, data)
}

