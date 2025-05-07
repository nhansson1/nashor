package handler

import (
	"github.com/gin-gonic/gin"
)

func (h Handler) HandleServePage(c *gin.Context) {
	c.File("./internal/dist/index.html")
}
