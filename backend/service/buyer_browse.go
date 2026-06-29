package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jesee-kuya/marineshop/domain"
)

func (b *Buyer) GetRecommendations(ctx context.Context, buyerID uuid.UUID) ([]*domain.Product, error) {
	return b.ProductRepo.GetRecommendations(ctx, buyerID)
}

func (b *Buyer) GetProducts(ctx context.Context, category string) ([]*domain.Product, error) {
	return b.ProductRepo.GetAllProducts(ctx, category)
}

func (b *Buyer) GetProductByID(ctx context.Context, productID uuid.UUID) (*domain.Product, error) {
	product, err := b.ProductRepo.FindProductByID(ctx, productID)
	if err != nil {
		return nil, err
	}
	if product == nil || !product.IsActive {
		return nil, domain.ErrProductNotFound
	}
	return product, nil
}
