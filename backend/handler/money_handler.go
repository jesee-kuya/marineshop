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

func (shop *Marineshop) GetMyPaymentAccounts(c *gin.Context) {
	claims, ok := c.MustGet(middleware.ClaimsKey).(*domain.Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	accounts, err := shop.MoneyService.GetMyPaymentAccounts(c.Request.Context(), claims.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch payment accounts"})
		return
	}

	c.JSON(http.StatusOK, accounts)
}

func (shop *Marineshop) Withdraw(c *gin.Context) {
	claims, ok := c.MustGet(middleware.ClaimsKey).(*domain.Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var req domain.WithdrawRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx, err := shop.MoneyService.Withdraw(c.Request.Context(), claims.UserID, &req)
	if err != nil {
		if errors.Is(err, domain.ErrInsufficientBalance) {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to process withdrawal"})
		return
	}

	c.JSON(http.StatusCreated, tx)
}

func (shop *Marineshop) TransactionHistory(c *gin.Context) {
	claims, ok := c.MustGet(middleware.ClaimsKey).(*domain.Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	transactions, err := shop.MoneyService.GetTransactionHistory(c.Request.Context(), claims.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch transaction history"})
		return
	}

	c.JSON(http.StatusOK, transactions)
}
