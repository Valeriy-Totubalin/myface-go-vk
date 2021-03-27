package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	return router
}
