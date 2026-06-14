package service

import (
	"github.com/jesee-kuya/marineshop/domain"
	"github.com/jesee-kuya/marineshop/repository"
)

type Auth struct {
	UserRepo repository.UserRepository
	JwtCfg   *domain.JWTConfig
}
