package handler

import "github.com/gin-gonic/gin"

type AuthHandler interface {
	Login(c *gin.Context)
	Signup(c *gin.Context)
	Logout(c *gin.Context)
	ChangePassword(c *gin.Context)
	ResetPassword(c *gin.Context)
}

type SellerHandler interface {
	CollectkYC(c *gin.Context)
	SetUpShop(c *gin.Context)
	SetUpPayment(c *gin.Context)
	CreateProduct(c *gin.Context)
	UpdateProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
	GetMyProducts(c *gin.Context)
	Withdraw(c *gin.Context)
	TransactionHistory(c *gin.Context)
	OrderManagement(c *gin.Context)
	Analytics(c *gin.Context)
	Profile(c *gin.Context)
}

type BuyerHandler interface {
	Recomendations(c *gin.Context)
	GetProducts(c *gin.Context)
	Profile(c *gin.Context)
	AddCartItem(c *gin.Context)
	GetCartItems(c *gin.Context)
	RemoveCartItem(c *gin.Context)
	Checkout(c *gin.Context)
	TransactionHistory(c *gin.Context)
	OrderManagement(c *gin.Context)
}
type AdminHandler interface {
	GETKYC(c *gin.Context)
	ApproveKYC(c *gin.Context)
	RejectKYC(c *gin.Context)
}
