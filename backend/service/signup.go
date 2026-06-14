package service

import (
	"context"

	"github.com/jesee-kuya/marineshop/domain"
	"github.com/jesee-kuya/marineshop/utils"
)

func (s *Auth) Signup(ctx context.Context, req *domain.SignupRequest) (*domain.AuthResponse, error) {
	existing, _ := s.UserRepo.FindByEmail(ctx, req.Email)
	if existing != nil {
		return nil, domain.ErrEmailInUse
	}

	hashed, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashed,
		Role:     "user",
	}

	created, err := s.UserRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	token, err := utils.GenerateToken(created, s.JwtCfg)
	if err != nil {
		return nil, err
	}

	return &domain.AuthResponse{Token: token, User: *created}, nil
}
