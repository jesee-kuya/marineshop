package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jesee-kuya/marineshop/domain"
)

func (m *Money) Withdraw(ctx context.Context, sellerID uuid.UUID, req *domain.WithdrawRequest) (*domain.SellerTransaction, error) {
	account, err := m.MoneyRepo.FindPaymentAccountByID(ctx, req.PaymentAccountID)
	if err != nil {
		return nil, err
	}
	if account == nil || account.SellerID != sellerID {
		return nil, domain.ErrUnauthorized
	}

	balance, err := m.MoneyRepo.GetSellerBalance(ctx, sellerID)
	if err != nil {
		return nil, err
	}
	if req.Amount > balance {
		return nil, domain.ErrInsufficientBalance
	}

	tx := &domain.SellerTransaction{
		SellerID:         sellerID,
		Type:             "withdrawal",
		Amount:           req.Amount,
		Status:           "pending",
		Description:      "withdrawal request",
		PaymentAccountID: &req.PaymentAccountID,
	}
	return m.MoneyRepo.CreateTransaction(ctx, tx)
}

func (m *Money) GetMyPaymentAccounts(ctx context.Context, sellerID uuid.UUID) ([]*domain.SellerPaymentAccount, error) {
	return m.MoneyRepo.FindPaymentAccountsBySellerID(ctx, sellerID)
}

func (m *Money) GetTransactionHistory(ctx context.Context, sellerID uuid.UUID) ([]*domain.SellerTransaction, error) {
	return m.MoneyRepo.GetTransactionsBySellerID(ctx, sellerID)
}
