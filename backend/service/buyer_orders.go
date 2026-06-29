package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jesee-kuya/marineshop/domain"
)

func (b *Buyer) GetOrders(ctx context.Context, buyerID uuid.UUID) ([]*domain.Order, error) {
	return b.OrderRepo.GetOrdersByBuyerID(ctx, buyerID)
}

func (b *Buyer) GetTransactionHistory(ctx context.Context, buyerID uuid.UUID) ([]*domain.BuyerTransaction, error) {
	return b.BuyerRepo.GetBuyerTransactionsByBuyerID(ctx, buyerID)
}
