package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jesee-kuya/marineshop/domain"
	"github.com/jesee-kuya/marineshop/repository"
)

func (b *Buyer) Checkout(ctx context.Context, buyerID uuid.UUID, req *domain.CheckoutRequest) ([]*domain.Order, error) {
	items, err := b.BuyerRepo.GetCartItemsByBuyerID(ctx, buyerID)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, domain.ErrCartEmpty
	}

	for _, item := range items {
		if item.Product.Stock < item.Quantity {
			return nil, domain.ErrOutOfStock
		}
	}

	var orders []*domain.Order

	err = b.Store.ExecCheckoutTx(ctx, func(
		buyerRepo repository.BuyerRepository,
		productRepo repository.ProductRepository,
		orderRepo repository.OrderRepository,
		moneyRepo repository.MoneyRepository,
	) error {
		for _, item := range items {
			order := &domain.Order{
				BuyerID:   buyerID,
				SellerID:  item.Product.SellerID,
				ProductID: item.ProductID,
				Quantity:  item.Quantity,
				Total:     item.Product.Price * float64(item.Quantity),
				Status:    "pending",
			}

			created, err := orderRepo.CreateOrder(ctx, order)
			if err != nil {
				return err
			}

			if err := productRepo.DeductStock(ctx, item.ProductID, item.Quantity); err != nil {
				return err
			}

			buyerTx := &domain.BuyerTransaction{
				BuyerID:   buyerID,
				OrderID:   created.ID,
				Amount:    created.Total,
				Status:    "pending",
				Reference: req.PaymentMethod,
			}
			if _, err := buyerRepo.CreateBuyerTransaction(ctx, buyerTx); err != nil {
				return err
			}

			sellerTx := &domain.SellerTransaction{
				SellerID:    item.Product.SellerID,
				Type:        "credit",
				Amount:      created.Total,
				Status:      "pending",
				Description: "sale of " + item.Product.Name,
			}
			if _, err := moneyRepo.CreateTransaction(ctx, sellerTx); err != nil {
				return err
			}

			orders = append(orders, created)
		}

		return buyerRepo.ClearCart(ctx, buyerID)
	})
	if err != nil {
		return nil, err
	}

	return orders, nil
}
