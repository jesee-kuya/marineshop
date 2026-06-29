package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jesee-kuya/marineshop/domain"
)

type AuthService interface {
	Signup(ctx context.Context, req *domain.SignupRequest) (*domain.AuthResponse, error)
	Login(ctx context.Context, req *domain.LoginRequest) (*domain.AuthResponse, error)
	ChangePassword(ctx context.Context, userID uuid.UUID, req *domain.ChangePasswordRequest) error
	ResetPassword(ctx context.Context, req *domain.ResetPasswordRequest) error
}

type SellerService interface {
	CollectKYC(ctx context.Context, userID uuid.UUID, req *domain.CollectKYCRequest) (*domain.SellerKYC, error)
	SetUpShop(ctx context.Context, userID uuid.UUID, req *domain.SetUpShopRequest) (*domain.BusinessKYC, error)
	GetProfile(ctx context.Context, userID uuid.UUID) (*domain.SellerProfile, error)
	GetOrders(ctx context.Context, sellerID uuid.UUID) ([]*domain.Order, error)
	UpdateOrderStatus(ctx context.Context, sellerID uuid.UUID, orderID uuid.UUID, status string) (*domain.Order, error)
	GetAnalytics(ctx context.Context, sellerID uuid.UUID) (*domain.SellerAnalytics, error)
}

type MoneyService interface {
	SetUpPayment(ctx context.Context, userID uuid.UUID, req *domain.SetUpPaymentRequest) (*domain.SellerPaymentAccount, error)
	GetMyPaymentAccounts(ctx context.Context, sellerID uuid.UUID) ([]*domain.SellerPaymentAccount, error)
	Withdraw(ctx context.Context, sellerID uuid.UUID, req *domain.WithdrawRequest) (*domain.SellerTransaction, error)
	GetTransactionHistory(ctx context.Context, sellerID uuid.UUID) ([]*domain.SellerTransaction, error)
}

type ProductService interface {
	CreateProduct(ctx context.Context, sellerID uuid.UUID, req *domain.CreateProductRequest) (*domain.Product, error)
	UpdateProduct(ctx context.Context, sellerID uuid.UUID, productID uuid.UUID, req *domain.UpdateProductRequest) (*domain.Product, error)
	DeleteProduct(ctx context.Context, sellerID uuid.UUID, productID uuid.UUID) error
	GetMyProducts(ctx context.Context, sellerID uuid.UUID) ([]*domain.Product, error)
}

type AdminService interface {
	GetPendingKYCs(ctx context.Context) ([]*domain.SellerKYC, error)
	ApproveKYC(ctx context.Context, kycID uuid.UUID) (*domain.SellerKYC, error)
	RejectKYC(ctx context.Context, kycID uuid.UUID) (*domain.SellerKYC, error)
}

type BuyerService interface {
	GetRecommendations(ctx context.Context, buyerID uuid.UUID) ([]*domain.Product, error)
	GetProducts(ctx context.Context, category string) ([]*domain.Product, error)
	GetProductByID(ctx context.Context, productID uuid.UUID) (*domain.Product, error)
	GetProfile(ctx context.Context, buyerID uuid.UUID) (*domain.User, error)
	AddCartItem(ctx context.Context, buyerID uuid.UUID, req *domain.AddCartItemRequest) (*domain.CartItem, error)
	GetCartItems(ctx context.Context, buyerID uuid.UUID) ([]*domain.CartItem, error)
	RemoveCartItem(ctx context.Context, buyerID uuid.UUID, itemID uuid.UUID) error
	Checkout(ctx context.Context, buyerID uuid.UUID, req *domain.CheckoutRequest) ([]*domain.Order, error)
	GetTransactionHistory(ctx context.Context, buyerID uuid.UUID) ([]*domain.BuyerTransaction, error)
	GetOrders(ctx context.Context, buyerID uuid.UUID) ([]*domain.Order, error)
}
