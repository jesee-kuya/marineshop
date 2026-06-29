package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jesee-kuya/marineshop/domain"
)

func (b *Buyer) AddCartItem(ctx context.Context, buyerID uuid.UUID, req *domain.AddCartItemRequest) (*domain.CartItem, error) {
	product, err := b.ProductRepo.FindProductByID(ctx, req.ProductID)
	if err != nil {
		return nil, err
	}
	if product == nil || !product.IsActive {
		return nil, domain.ErrProductNotFound
	}
	if product.Stock < req.Quantity {
		return nil, domain.ErrOutOfStock
	}

	existing, err := b.BuyerRepo.FindCartItemByProductID(ctx, buyerID, req.ProductID)
	if err != nil {
		return nil, err
	}

	if existing != nil {
		newQty := existing.Quantity + req.Quantity
		if product.Stock < newQty {
			return nil, domain.ErrOutOfStock
		}
		return b.BuyerRepo.UpdateCartItemQuantity(ctx, existing.ID, newQty)
	}

	item := &domain.CartItem{
		BuyerID:   buyerID,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}
	return b.BuyerRepo.InsertCartItem(ctx, item)
}

func (b *Buyer) GetCartItems(ctx context.Context, buyerID uuid.UUID) ([]*domain.CartItem, error) {
	return b.BuyerRepo.GetCartItemsByBuyerID(ctx, buyerID)
}

func (b *Buyer) RemoveCartItem(ctx context.Context, buyerID uuid.UUID, itemID uuid.UUID) error {
	return b.BuyerRepo.RemoveCartItem(ctx, itemID, buyerID)
}
