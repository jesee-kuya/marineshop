package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jesee-kuya/marineshop/domain"
)

func (h *Authentication) Signup(c *gin.Context) {
	var req domain.SignupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.AuthService.Signup(c.Request.Context(), &req)
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
