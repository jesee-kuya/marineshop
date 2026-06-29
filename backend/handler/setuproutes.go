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

	api.GET("/products", shop.GetProducts)
	api.GET("/products/:id", shop.GetProductByID)

	protected := api.Group("/")
	protected.Use(shop.Middleware.AuthMiddleware())
	{
		protected.POST("/auth/change-password", shop.ChangePassword)
		protected.POST("/auth/logout", shop.Logout)

		buyerGroup := protected.Group("/buyer")
		buyerGroup.Use(shop.Middleware.RequireRole("buyer"))
		{
			buyerGroup.GET("/recommendations", shop.Recomendations)
			buyerGroup.GET("/profile", shop.BuyerProfile)
			buyerGroup.POST("/cart", shop.AddCartItem)
			buyerGroup.GET("/cart", shop.GetCartItems)
			buyerGroup.DELETE("/cart/:id", shop.RemoveCartItem)
			buyerGroup.POST("/checkout", shop.Checkout)
			buyerGroup.GET("/transactions", shop.BuyerTransactionHistory)
			buyerGroup.GET("/orders", shop.BuyerOrderManagement)
		}

		adminGroup := protected.Group("/admin")
		adminGroup.Use(shop.Middleware.RequireRole("admin"))
		{
			adminGroup.GET("/kyc", shop.GETKYC)
			adminGroup.PUT("/kyc/:id/approve", shop.ApproveKYC)
			adminGroup.PUT("/kyc/:id/reject", shop.RejectKYC)
		}

		sellerGroup := protected.Group("/seller")
		sellerGroup.Use(shop.Middleware.RequireRole("seller"))
		{
			sellerGroup.POST("/kyc", shop.CollectkYC)
			sellerGroup.POST("/shop", shop.SetUpShop)
			sellerGroup.POST("/payment", shop.SetUpPayment)
			sellerGroup.GET("/payments", shop.GetMyPaymentAccounts)
			sellerGroup.POST("/products", shop.CreateProduct)
			sellerGroup.GET("/products", shop.GetMyProducts)
			sellerGroup.PUT("/products/:id", shop.UpdateProduct)
			sellerGroup.DELETE("/products/:id", shop.DeleteProduct)
			sellerGroup.POST("/withdraw", shop.Withdraw)
			sellerGroup.GET("/transactions", shop.TransactionHistory)
			sellerGroup.GET("/orders", shop.OrderManagement)
			sellerGroup.PUT("/orders/:id", shop.UpdateOrderStatus)
			sellerGroup.GET("/analytics", shop.Analytics)
			sellerGroup.GET("/profile", shop.Profile)
		}
	}

	return router
}
