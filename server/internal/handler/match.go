package handler

import (
	"github.com/gin-gonic/gin"
	"nashor/internal/problem"
)

func (h Handler) HandleGetMatchesByPuuid(c *gin.Context) {
	var (
		server = c.Param("server")
		puuid  = c.Param("puuid")
		start  = c.Param("start")
		count  = c.Param("count")
	)

	ids, err := h.matchService.GetMatchIdsByPuuid(server, puuid, start, count)

	if err != nil {
		var perr problem.ErrorResponse
		perr.FromErr(err)

		c.JSON(perr.Status, perr)
		return
	}

	matches, err := h.matchService.GetMatchDataByIds(server, ids)

	if err != nil {
		var perr problem.ErrorResponse
		perr.FromErr(err)

		c.JSON(perr.Status, perr)
		return
	}

	c.JSON(200, matches)
}
