package service

import (
	"github.com/jesee-kuya/marineshop/domain"
	"github.com/jesee-kuya/marineshop/repository"
)

func NewAuthService(userRepo repository.UserRepository, jwtCfg *domain.JWTConfig) AuthService {
	return &Auth{UserRepo: userRepo, JwtCfg: jwtCfg}
}

func NewSellerService(kycRepo repository.KYCRepository) SellerService {
	return &Seller{KYCRepo: kycRepo}
}

func NewMoneyService(moneyRepo repository.MoneyRepository, kycRepo repository.KYCRepository) MoneyService {
	return &Money{MoneyRepo: moneyRepo, KYCRepo: kycRepo}
}
