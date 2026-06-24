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
}
