package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jesee-kuya/marineshop/domain"
	"github.com/jesee-kuya/marineshop/middleware"
)

func (shop *Marineshop) Recomendations(c *gin.Context) {
	claims, ok := c.MustGet(middleware.ClaimsKey).(*domain.Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	products, err := shop.BuyerService.GetRecommendations(c.Request.Context(), claims.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch recommendations"})
		return
	}

	c.JSON(http.StatusOK, products)
}

func (shop *Marineshop) GetProducts(c *gin.Context) {
	category := c.Query("category")

	products, err := shop.BuyerService.GetProducts(c.Request.Context(), category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch products"})
		return
	}

	c.JSON(http.StatusOK, products)
}

func (shop *Marineshop) GetProductByID(c *gin.Context) {
	productID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}

	product, err := shop.BuyerService.GetProductByID(c.Request.Context(), productID)
	if err != nil {
		if errors.Is(err, domain.ErrProductNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch product"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (shop *Marineshop) AddCartItem(c *gin.Context) {
	claims, ok := c.MustGet(middleware.ClaimsKey).(*domain.Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var req domain.AddCartItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := shop.BuyerService.AddCartItem(c.Request.Context(), claims.UserID, &req)
	if err != nil {
		if errors.Is(err, domain.ErrProductNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if errors.Is(err, domain.ErrOutOfStock) {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to add item to cart"})
		return
	}

	c.JSON(http.StatusCreated, item)
}

func (shop *Marineshop) GetCartItems(c *gin.Context) {
	claims, ok := c.MustGet(middleware.ClaimsKey).(*domain.Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	items, err := shop.BuyerService.GetCartItems(c.Request.Context(), claims.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch cart"})
		return
	}

	c.JSON(http.StatusOK, items)
}

func (shop *Marineshop) RemoveCartItem(c *gin.Context) {
	claims, ok := c.MustGet(middleware.ClaimsKey).(*domain.Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	itemID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid cart item id"})
		return
	}

	if err := shop.BuyerService.RemoveCartItem(c.Request.Context(), claims.UserID, itemID); err != nil {
		if errors.Is(err, domain.ErrCartItemNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to remove cart item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "item removed from cart"})
}

func (shop *Marineshop) Checkout(c *gin.Context) {
	claims, ok := c.MustGet(middleware.ClaimsKey).(*domain.Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var req domain.CheckoutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orders, err := shop.BuyerService.Checkout(c.Request.Context(), claims.UserID, &req)
	if err != nil {
		if errors.Is(err, domain.ErrCartEmpty) {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}
		if errors.Is(err, domain.ErrOutOfStock) {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to process checkout"})
		return
	}

	c.JSON(http.StatusCreated, orders)
}

func (shop *Marineshop) BuyerProfile(c *gin.Context) {
	claims, ok := c.MustGet(middleware.ClaimsKey).(*domain.Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	profile, err := shop.BuyerService.GetProfile(c.Request.Context(), claims.UserID)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch profile"})
		return
	}

	c.JSON(http.StatusOK, profile)
}

func (shop *Marineshop) BuyerTransactionHistory(c *gin.Context) {
	claims, ok := c.MustGet(middleware.ClaimsKey).(*domain.Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	transactions, err := shop.BuyerService.GetTransactionHistory(c.Request.Context(), claims.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch transaction history"})
		return
	}

	c.JSON(http.StatusOK, transactions)
}

func (shop *Marineshop) BuyerOrderManagement(c *gin.Context) {
	claims, ok := c.MustGet(middleware.ClaimsKey).(*domain.Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	orders, err := shop.BuyerService.GetOrders(c.Request.Context(), claims.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch orders"})
		return
	}

	c.JSON(http.StatusOK, orders)
}
