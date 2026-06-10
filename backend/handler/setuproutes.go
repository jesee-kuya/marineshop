package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jesee-kuya/marineshop/middleware"
)

func SetupRoutes(router *gin.Engine, mw middleware.Middleware, authHandler AuthHandler) {
	router.Use(mw.RouteChecker())

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	api := router.Group("/api/v1")

	authGroup := api.Group("/auth")
	{
		authGroup.POST("/signup", authHandler.Signup)
		authGroup.POST("/login", authHandler.Login)
	}
}
