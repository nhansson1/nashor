package handler

import (
	"nashor/internal/services"
)

type Handler struct {
	accountService  *services.AccountService
	summonerService *services.SummonerService
	matchService    *services.MatchService
	leagueService   *services.LeagueService
}

func NewHandler(a *services.AccountService, s *services.SummonerService, m *services.MatchService, l *services.LeagueService) Handler {
	return Handler{
		accountService:  a,
		summonerService: s,
		matchService:    m,
		leagueService:   l,
	}
}
