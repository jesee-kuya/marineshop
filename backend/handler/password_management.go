package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (shop *Marineshop) ChangePassword(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "not implemented"})
}

func (shop *Marineshop) ResetPassword(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "not implemented"})
}
