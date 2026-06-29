package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jesee-kuya/marineshop/domain"
)

func (shop *Marineshop) Signup(c *gin.Context) {
	var req domain.SignupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Role == "admin" {
		if shop.AdminSecret == "" || req.AdminSecret != shop.AdminSecret {
			c.JSON(http.StatusForbidden, gin.H{"error": domain.ErrInvalidAdminSecret.Error()})
			return
		}
	}

	res, err := shop.AuthService.Signup(c.Request.Context(), &req)
	if err != nil {
		if errors.Is(err, domain.ErrEmailInUse) {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create account"})
		return
	}

	c.JSON(http.StatusCreated, res)
}
