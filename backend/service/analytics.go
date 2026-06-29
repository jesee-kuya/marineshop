package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jesee-kuya/marineshop/domain"
)

func (s *Seller) GetAnalytics(ctx context.Context, sellerID uuid.UUID) (*domain.SellerAnalytics, error) {
	productCount, err := s.ProductRepo.CountProductsBySellerID(ctx, sellerID)
	if err != nil {
		return nil, err
	}

	total, pending, completed, err := s.OrderRepo.CountOrdersBySellerID(ctx, sellerID)
	if err != nil {
		return nil, err
	}

	revenue, err := s.OrderRepo.GetTotalRevenueBySellerID(ctx, sellerID)
	if err != nil {
		return nil, err
	}

	return &domain.SellerAnalytics{
		TotalProducts:   productCount,
		TotalOrders:     total,
		TotalRevenue:    revenue,
		PendingOrders:   pending,
		CompletedOrders: completed,
	}, nil
}
