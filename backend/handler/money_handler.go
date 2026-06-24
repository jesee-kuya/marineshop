package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jesee-kuya/marineshop/domain"
	"github.com/jesee-kuya/marineshop/middleware"
)

func (shop *Marineshop) SetUpPayment(c *gin.Context) {
	claims, ok := c.MustGet(middleware.ClaimsKey).(*domain.Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var req domain.SetUpPaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	account, err := shop.MoneyService.SetUpPayment(c.Request.Context(), claims.UserID, &req)
	if err != nil {
		if errors.Is(err, domain.ErrSellerKYCNotFound) {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}
		if errors.Is(err, domain.ErrInvalidPaymentDetails) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to set up payment account"})
		return
	}

	c.JSON(http.StatusCreated, account)
}
