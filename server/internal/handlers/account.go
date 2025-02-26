package handlers

import (
	"errors"
	"fmt"
	"nashor/internal/problem"
	"nashor/internal/services"
	"github.com/gin-gonic/gin"
)

func AccountByRiotIdHandler(c *gin.Context) {
    gameName := c.Param("gameName")
    tagLine := c.Param("tagLine")

    data, err := services.GetAccountByRiotId(gameName, tagLine)

    var perr problem.ErrorResponse
    if err != nil {

        if errors.As(err, &perr) {
            c.JSON(perr.Status, perr)
            return
        }

        fmt.Println(err)
        c.JSON(500, problem.InternalServerError())
        return
    }

    c.JSON(200, data)
}
