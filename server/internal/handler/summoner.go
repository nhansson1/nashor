package handler

import (
	"github.com/gin-gonic/gin"
	"nashor/internal/problem"
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

func (h Handler) HandleGetSummonerProfile(c *gin.Context) {
	var (
		region   = c.Param("region")
		gameName = c.Param("gameName")
		tagLine  = c.Param("tagLine")
	)
	acc, err := h.accountService.GetAccountByRiotId(gameName, tagLine)

	if err != nil {
		var res problem.ErrorResponse
		res.FromErr(err)

		c.JSON(res.Status, res)
		return
	}

	s, err := h.summonerService.GetSummonerByPuuid(region, acc.Puuid)

	if err != nil {
		var res problem.ErrorResponse
		res.FromErr(err)

		c.JSON(res.Status, res)
		return
	}

	entries, err := h.leagueService.GetRankQueusById(region, s.SummonerId)

	if err != nil {
		var res problem.ErrorResponse
		res.FromErr(err)

		c.JSON(res.Status, res)
		return
	}

	c.JSON(200, gin.H{
		"summoner": s,
		"account": gin.H{
			"gameName": acc.GameName,
			"tagLine":  acc.TagLine,
		},
		"ranks": entries,
	})
}
