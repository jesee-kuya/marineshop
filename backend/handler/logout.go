package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Authentication) Logout(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "not implemented"})
}
