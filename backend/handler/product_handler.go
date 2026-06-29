package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jesee-kuya/marineshop/domain"
	"github.com/jesee-kuya/marineshop/middleware"
)

func (shop *Marineshop) CreateProduct(c *gin.Context) {
	claims, ok := c.MustGet(middleware.ClaimsKey).(*domain.Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var req domain.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := shop.ProductService.CreateProduct(c.Request.Context(), claims.UserID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create product"})
		return
	}

	c.JSON(http.StatusCreated, product)
}

func (shop *Marineshop) UpdateProduct(c *gin.Context) {
	claims, ok := c.MustGet(middleware.ClaimsKey).(*domain.Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	productID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}

	var req domain.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := shop.ProductService.UpdateProduct(c.Request.Context(), claims.UserID, productID, &req)
	if err != nil {
		if errors.Is(err, domain.ErrProductNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update product"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (shop *Marineshop) DeleteProduct(c *gin.Context) {
	claims, ok := c.MustGet(middleware.ClaimsKey).(*domain.Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	productID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}

	if err := shop.ProductService.DeleteProduct(c.Request.Context(), claims.UserID, productID); err != nil {
		if errors.Is(err, domain.ErrProductNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product deleted"})
}

func (shop *Marineshop) GetMyProducts(c *gin.Context) {
	claims, ok := c.MustGet(middleware.ClaimsKey).(*domain.Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	products, err := shop.ProductService.GetMyProducts(c.Request.Context(), claims.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch products"})
		return
	}

	c.JSON(http.StatusOK, products)
}
