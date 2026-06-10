package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Authentication) ChangePassword(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "not implemented"})
}

func (h *Authentication) ResetPassword(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "not implemented"})
}
