package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (shop *Marineshop) Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "logged out successfully"})
}
