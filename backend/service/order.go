package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jesee-kuya/marineshop/domain"
)

func (s *Seller) GetOrders(ctx context.Context, sellerID uuid.UUID) ([]*domain.Order, error) {
	return s.OrderRepo.GetOrdersBySellerID(ctx, sellerID)
}

func (s *Seller) UpdateOrderStatus(ctx context.Context, sellerID uuid.UUID, orderID uuid.UUID, status string) (*domain.Order, error) {
	return s.OrderRepo.UpdateOrderStatus(ctx, sellerID, orderID, status)
}
