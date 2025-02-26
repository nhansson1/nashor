package handlers

import (
    "errors"
    "nashor/internal/services"
    "nashor/internal/problem"
    "github.com/gin-gonic/gin"
)

func SummonerByPuuid(c *gin.Context) {
    puuid := c.Param("puuid")

    data, err := services.GetSummonerByPuuid(puuid)

    if err != nil {
        var perr problem.ErrorResponse
        if errors.As(err, &perr) {
            c.JSON(perr.Status, perr)
            return
        }

        c.JSON(500, problem.InternalServerError())
        return
    }

    c.JSON(200, data)
}
