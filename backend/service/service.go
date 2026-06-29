package service

import (
	"github.com/jesee-kuya/marineshop/domain"
	"github.com/jesee-kuya/marineshop/repository"
)

func NewAuthService(userRepo repository.UserRepository, jwtCfg *domain.JWTConfig) AuthService {
	return &Auth{UserRepo: userRepo, JwtCfg: jwtCfg}
}

func NewSellerService(kycRepo repository.KYCRepository, productRepo repository.ProductRepository, orderRepo repository.OrderRepository) SellerService {
	return &Seller{KYCRepo: kycRepo, ProductRepo: productRepo, OrderRepo: orderRepo}
}

func NewMoneyService(moneyRepo repository.MoneyRepository, kycRepo repository.KYCRepository) MoneyService {
	return &Money{MoneyRepo: moneyRepo, KYCRepo: kycRepo}
}

func NewProductService(productRepo repository.ProductRepository) ProductService {
	return &ProductSvc{ProductRepo: productRepo}
}

func NewBuyerService(buyerRepo repository.BuyerRepository, productRepo repository.ProductRepository, orderRepo repository.OrderRepository, userRepo repository.UserRepository, store repository.CheckoutStore) BuyerService {
	return &Buyer{BuyerRepo: buyerRepo, ProductRepo: productRepo, OrderRepo: orderRepo, UserRepo: userRepo, Store: store}
}

func NewAdminService(kycRepo repository.KYCRepository) AdminService {
	return &Admin{KYCRepo: kycRepo}
}
