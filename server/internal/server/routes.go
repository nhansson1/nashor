package server

import (
	"nashor/internal/handler"

	"github.com/gin-gonic/gin"
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
			s.GET("/profile/by-riot-id/:region/:gameName/:tagLine", h.HandleGetSummonerProfile)
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
