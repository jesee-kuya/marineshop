package service

import (
	"context"

	"github.com/jesee-kuya/marineshop/domain"
	"github.com/jesee-kuya/marineshop/utils"
)

func (s *Auth) Login(ctx context.Context, req *domain.LoginRequest) (*domain.AuthResponse, error) {
	user, err := s.UserRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, domain.ErrInvalidCredentials
	}

	if !utils.CheckPassword(req.Password, user.Password) {
		return nil, domain.ErrInvalidCredentials
	}

	token, err := utils.GenerateToken(user, s.JwtCfg)
	if err != nil {
		return nil, err
	}

	user.Password = ""
	return &domain.AuthResponse{Token: token, User: *user}, nil
}
