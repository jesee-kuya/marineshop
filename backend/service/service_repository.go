package service

import (
	"context"

	"github.com/jesee-kuya/marineshop/domain"
)

type AuthService interface {
	Signup(ctx context.Context, req *domain.SignupRequest) (*domain.AuthResponse, error)
}
