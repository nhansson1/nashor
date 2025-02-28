package server

import (
	"github.com/gin-gonic/gin"
	"nashor/internal/handlers"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		a := api.Group("/account")
		{
			a.GET("/by-riot-id/:gameName/:tagLine", handlers.HandleGetAccountByRiotID)
		}

		s := api.Group("/summoner")
		{
			s.GET("/by-puuid/:region/:puuid", handlers.HandleGetSummonerById)
		}

		l := api.Group("/league")
		{
			l.GET("/entries/by-id/:region/:summonerId", handlers.HandleGetLeagueEntriesById)
		}
		m := api.Group("/matches")
		{
			m.GET("by-puuid/:puuid/:start/:count", handlers.HandleGetMatchesByPuuid)
		}
	}
}
