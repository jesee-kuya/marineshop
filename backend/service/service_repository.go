package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jesee-kuya/marineshop/domain"
)

type AuthService interface {
	Signup(ctx context.Context, req *domain.SignupRequest) (*domain.AuthResponse, error)
	Login(ctx context.Context, req *domain.LoginRequest) (*domain.AuthResponse, error)
	ChangePassword(ctx context.Context, userID uuid.UUID, req *domain.ChangePasswordRequest) error
	ResetPassword(ctx context.Context, req *domain.ResetPasswordRequest) error
}

type SellerService interface {
	CollectKYC(ctx context.Context, userID uuid.UUID, req *domain.CollectKYCRequest) (*domain.SellerKYC, error)
	SetUpShop(ctx context.Context, userID uuid.UUID, req *domain.SetUpShopRequest) (*domain.BusinessKYC, error)
}
