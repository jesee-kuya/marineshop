package handler

import (
	"github.com/jesee-kuya/marineshop/middleware"
	"github.com/jesee-kuya/marineshop/service"
)

type Marineshop struct {
	AuthService    service.AuthService
	SellerService  service.SellerService
	MoneyService   service.MoneyService
	ProductService service.ProductService
	BuyerService   service.BuyerService
	AdminService   service.AdminService
	Middleware     middleware.Middleware
	AdminSecret    string
}
