package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jesee-kuya/marineshop/domain"
)

func (p *ProductSvc) CreateProduct(ctx context.Context, sellerID uuid.UUID, req *domain.CreateProductRequest) (*domain.Product, error) {
	product := &domain.Product{
		SellerID:    sellerID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		Category:    req.Category,
		ImageURL:    req.ImageURL,
	}
	return p.ProductRepo.CreateProduct(ctx, product)
}

func (p *ProductSvc) UpdateProduct(ctx context.Context, sellerID uuid.UUID, productID uuid.UUID, req *domain.UpdateProductRequest) (*domain.Product, error) {
	return p.ProductRepo.UpdateProduct(ctx, productID, sellerID, req)
}

func (p *ProductSvc) DeleteProduct(ctx context.Context, sellerID uuid.UUID, productID uuid.UUID) error {
	return p.ProductRepo.DeleteProduct(ctx, productID, sellerID)
}

func (p *ProductSvc) GetMyProducts(ctx context.Context, sellerID uuid.UUID) ([]*domain.Product, error) {
	return p.ProductRepo.GetProductsBySellerID(ctx, sellerID)
}
