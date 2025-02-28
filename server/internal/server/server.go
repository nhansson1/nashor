package server

import "github.com/gin-gonic/gin"

func Run() {
	router := gin.New()

	RegisterRoutes(router)

	router.Run()
}
