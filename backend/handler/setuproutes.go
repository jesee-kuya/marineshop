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
		authGroup.POST("/signup", shop.Signup)
		authGroup.POST("/login", shop.Login)
		authGroup.POST("/reset-password", shop.ResetPassword)
	}

	protected := api.Group("/")
	protected.Use(shop.Middleware.AuthMiddleware())
	{
		protected.POST("/auth/change-password", shop.ChangePassword)

		sellerGroup := protected.Group("/seller")
		{
			sellerGroup.POST("/kyc", shop.CollectkYC)
			sellerGroup.POST("/shop", shop.SetUpShop)
		}
	}

	return router
}
