package server

import (
	"nashor/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
    api := r.Group("/api/v1")
    {
        acc := api.Group("/account")
        {
            acc.GET("/by-riot-id/:gameName/:tagLine", handlers.AccountByRiotIdHandler)
        }

        sum := api.Group("/summoner")
        {
            sum.GET("/by-puuid/:puuid", handlers.SummonerByPuuid)
        }
    }
}
