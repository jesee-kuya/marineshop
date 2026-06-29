package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jesee-kuya/marineshop/domain"
)

func (shop *Marineshop) GETKYC(c *gin.Context) {
	kycs, err := shop.AdminService.GetPendingKYCs(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch pending KYCs"})
		return
	}

	c.JSON(http.StatusOK, kycs)
}

func (shop *Marineshop) ApproveKYC(c *gin.Context) {
	kycID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid kyc id"})
		return
	}

	kyc, err := shop.AdminService.ApproveKYC(c.Request.Context(), kycID)
	if err != nil {
		if errors.Is(err, domain.ErrKYCNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to approve KYC"})
		return
	}

	c.JSON(http.StatusOK, kyc)
}

func (shop *Marineshop) RejectKYC(c *gin.Context) {
	kycID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid kyc id"})
		return
	}

	kyc, err := shop.AdminService.RejectKYC(c.Request.Context(), kycID)
	if err != nil {
		if errors.Is(err, domain.ErrKYCNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to reject KYC"})
		return
	}

	c.JSON(http.StatusOK, kyc)
}
