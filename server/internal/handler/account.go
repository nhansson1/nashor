package handler

import (
	"github.com/gin-gonic/gin"
	"nashor/internal/problem"
)

func (h Handler) HandleGetAccountByRiotId(c *gin.Context) {
	var (
		gameName = c.Param("gameName")
		tagLine  = c.Param("tagLine")
	)

	data, err := h.accountService.GetAccountByRiotId(gameName, tagLine)

	if err != nil {
		var res problem.ErrorResponse
		res.FromErr(err)

		c.JSON(res.Status, res)
		return
	}

	c.JSON(200, data)
}
