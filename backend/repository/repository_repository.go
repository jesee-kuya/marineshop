package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jesee-kuya/marineshop/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	FindByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
	UpdatePassword(ctx context.Context, id uuid.UUID, hashedPassword string) error
}

type KYCRepository interface {
	CreateSellerKYC(ctx context.Context, kyc *domain.SellerKYC) (*domain.SellerKYC, error)
	FindKYCByUserID(ctx context.Context, userID uuid.UUID) (*domain.SellerKYC, error)
	CreateBusinessKYC(ctx context.Context, kyc *domain.BusinessKYC) (*domain.BusinessKYC, error)
	FindBusinessKYCBySellerKYCID(ctx context.Context, sellerKYCID uuid.UUID) (*domain.BusinessKYC, error)
}

type MoneyRepository interface {
	CreateSellerPaymentAccount(ctx context.Context, account *domain.SellerPaymentAccount) (*domain.SellerPaymentAccount, error)
	FindPaymentAccountsBySellerID(ctx context.Context, sellerID uuid.UUID) ([]*domain.SellerPaymentAccount, error)
}
