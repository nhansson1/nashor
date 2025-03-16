package server

import (
	"nashor/internal/handler"
	"nashor/internal/services"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.New()

	cors := cors.Default()
	
	router.Use(cors)

    rc := services.NewRiotClient(os.Getenv("API_KEY"), os.Getenv("REGION"))

    accountService := services.NewAccountService(rc)
    summonerService := services.NewSummonerService(rc)
    matchService := services.NewMatchService(rc)
    leagueService := services.NewLeagueService(rc)

    h := handler.NewHandler(&accountService, &summonerService, &matchService, &leagueService)
	RegisterRoutes(router, &h)

	router.Run()
}
