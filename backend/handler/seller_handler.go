package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jesee-kuya/marineshop/domain"
	"github.com/jesee-kuya/marineshop/middleware"
)

func (shop *Marineshop) CollectkYC(c *gin.Context) {
	claims, ok := c.MustGet(middleware.ClaimsKey).(*domain.Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var req domain.CollectKYCRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	kyc, err := shop.SellerService.CollectKYC(c.Request.Context(), claims.UserID, &req)
	if err != nil {
		if errors.Is(err, domain.ErrKYCAlreadyExists) {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to submit KYC"})
		return
	}

	c.JSON(http.StatusCreated, kyc)
}

func (shop *Marineshop) SetUpShop(c *gin.Context) {
	claims, ok := c.MustGet(middleware.ClaimsKey).(*domain.Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var req domain.SetUpShopRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	businessKYC, err := shop.SellerService.SetUpShop(c.Request.Context(), claims.UserID, &req)
	if err != nil {
		if errors.Is(err, domain.ErrSellerKYCNotFound) {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}
		if errors.Is(err, domain.ErrBusinessKYCAlreadyExists) {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to set up shop"})
		return
	}

	c.JSON(http.StatusCreated, businessKYC)
}

func (shop *Marineshop) OrderManagement(c *gin.Context) {
	claims, ok := c.MustGet(middleware.ClaimsKey).(*domain.Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	orders, err := shop.SellerService.GetOrders(c.Request.Context(), claims.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch orders"})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func (shop *Marineshop) Analytics(c *gin.Context) {
	claims, ok := c.MustGet(middleware.ClaimsKey).(*domain.Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	analytics, err := shop.SellerService.GetAnalytics(c.Request.Context(), claims.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch analytics"})
		return
	}

	c.JSON(http.StatusOK, analytics)
}

func (shop *Marineshop) UpdateOrderStatus(c *gin.Context) {
	claims, ok := c.MustGet(middleware.ClaimsKey).(*domain.Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	orderID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order id"})
		return
	}

	var req domain.UpdateOrderStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := shop.SellerService.UpdateOrderStatus(c.Request.Context(), claims.UserID, orderID, req.Status)
	if err != nil {
		if errors.Is(err, domain.ErrOrderNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update order status"})
		return
	}

	c.JSON(http.StatusOK, order)
}

func (shop *Marineshop) Profile(c *gin.Context) {
	claims, ok := c.MustGet(middleware.ClaimsKey).(*domain.Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	profile, err := shop.SellerService.GetProfile(c.Request.Context(), claims.UserID)
	if err != nil {
		if errors.Is(err, domain.ErrSellerKYCNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch profile"})
		return
	}

	c.JSON(http.StatusOK, profile)
}
