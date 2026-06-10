package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (shop *Marineshop) SetupRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(shop.Middleware.RouteChecker())

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	api := router.Group("/api/v1")

	authGroup := api.Group("/auth")
	{
		authGroup.POST("/signup", shop.Auth.Signup)
		authGroup.POST("/login", shop.Auth.Login)
	}

	return router
}
