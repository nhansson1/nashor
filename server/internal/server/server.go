package server

import "github.com/gin-gonic/gin"

func Run() {
    router := gin.Default()

    RegisterRoutes(router)

    router.Run()
}
