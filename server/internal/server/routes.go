package server

import (
	"github.com/gin-gonic/gin"
	"nashor/internal/handler"
)

func RegisterRoutes(r *gin.Engine, h *handler.Handler) {
	api := r.Group("/api/v1")
	{
		a := api.Group("/account")
		{
			a.GET("/by-riot-id/:gameName/:tagLine", h.HandleGetAccountByRiotId)
		}

		s := api.Group("/summoner")
		{
			s.GET("/by-puuid/:region/:puuid", h.HandleGetSummonerById)
		}

		l := api.Group("/league")
		{
			l.GET("/entries/by-id/:region/:summonerId", h.HandleGetLeagueEntriesById)
		}
		m := api.Group("/matches")
		{
            m.GET("by-puuid/:server/:puuid/:start/:count", h.HandleGetMatchesByPuuid)
		}
	}
}
