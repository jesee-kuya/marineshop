package service

import (
	"github.com/jesee-kuya/marineshop/domain"
	"github.com/jesee-kuya/marineshop/repository"
)

type Auth struct {
	UserRepo repository.UserRepository
	JwtCfg   *domain.JWTConfig
}

type Seller struct {
	KYCRepo     repository.KYCRepository
	ProductRepo repository.ProductRepository
	OrderRepo   repository.OrderRepository
}

type Money struct {
	MoneyRepo repository.MoneyRepository
	KYCRepo   repository.KYCRepository
}

type ProductSvc struct {
	ProductRepo repository.ProductRepository
}

type Admin struct {
	KYCRepo repository.KYCRepository
}

type Buyer struct {
	BuyerRepo   repository.BuyerRepository
	ProductRepo repository.ProductRepository
	OrderRepo   repository.OrderRepository
	UserRepo    repository.UserRepository
	Store       repository.CheckoutStore
}
