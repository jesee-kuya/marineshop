package service

import (
	"context"

	"github.com/jesee-kuya/marineshop/domain"
	"github.com/jesee-kuya/marineshop/repository"
	"github.com/jesee-kuya/marineshop/utils"
)

type AuthService interface {
	Signup(ctx context.Context, req *domain.SignupRequest) (*domain.AuthResponse, error)
}

type authService struct {
	userRepo repository.UserRepository
	jwtCfg   *domain.JWTConfig
}

func NewAuthService(userRepo repository.UserRepository, jwtCfg *domain.JWTConfig) AuthService {
	return &authService{userRepo: userRepo, jwtCfg: jwtCfg}
}

func (s *authService) Signup(ctx context.Context, req *domain.SignupRequest) (*domain.AuthResponse, error) {
	existing, _ := s.userRepo.FindByEmail(ctx, req.Email)
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

	created, err := s.userRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	token, err := utils.GenerateToken(created, s.jwtCfg)
	if err != nil {
		return nil, err
	}

	return &domain.AuthResponse{Token: token, User: *created}, nil
}
