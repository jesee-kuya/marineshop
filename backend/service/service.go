package service

import (
	"github.com/jesee-kuya/marineshop/domain"
	"github.com/jesee-kuya/marineshop/repository"
)

func NewAuthService(userRepo repository.UserRepository, jwtCfg *domain.JWTConfig) AuthService {
	return &Auth{UserRepo: userRepo, JwtCfg: jwtCfg}
}
