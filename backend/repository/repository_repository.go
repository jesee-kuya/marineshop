package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jesee-kuya/marineshop/domain"
)

// CheckoutStore runs all checkout operations atomically inside a single transaction.
// The callback receives tx-scoped repository instances; a rollback is issued on any error.
type CheckoutStore interface {
	ExecCheckoutTx(ctx context.Context, fn func(
		buyerRepo BuyerRepository,
		productRepo ProductRepository,
		orderRepo OrderRepository,
		moneyRepo MoneyRepository,
	) error) error
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	FindByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
	UpdatePassword(ctx context.Context, id uuid.UUID, hashedPassword string) error
}

type KYCRepository interface {
	CreateSellerKYC(ctx context.Context, kyc *domain.SellerKYC) (*domain.SellerKYC, error)
	FindKYCByUserID(ctx context.Context, userID uuid.UUID) (*domain.SellerKYC, error)
	CreateBusinessKYC(ctx context.Context, kyc *domain.BusinessKYC) (*domain.BusinessKYC, error)
	FindBusinessKYCBySellerKYCID(ctx context.Context, sellerKYCID uuid.UUID) (*domain.BusinessKYC, error)
	GetPendingKYCs(ctx context.Context) ([]*domain.SellerKYC, error)
	UpdateKYCStatus(ctx context.Context, id uuid.UUID, status string) (*domain.SellerKYC, error)
}

type MoneyRepository interface {
	CreateSellerPaymentAccount(ctx context.Context, account *domain.SellerPaymentAccount) (*domain.SellerPaymentAccount, error)
	FindPaymentAccountsBySellerID(ctx context.Context, sellerID uuid.UUID) ([]*domain.SellerPaymentAccount, error)
	FindPaymentAccountByID(ctx context.Context, id uuid.UUID) (*domain.SellerPaymentAccount, error)
	CreateTransaction(ctx context.Context, tx *domain.SellerTransaction) (*domain.SellerTransaction, error)
	GetTransactionsBySellerID(ctx context.Context, sellerID uuid.UUID) ([]*domain.SellerTransaction, error)
	GetSellerBalance(ctx context.Context, sellerID uuid.UUID) (float64, error)
}

type ProductRepository interface {
	CreateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error)
	UpdateProduct(ctx context.Context, id uuid.UUID, sellerID uuid.UUID, req *domain.UpdateProductRequest) (*domain.Product, error)
	DeleteProduct(ctx context.Context, id uuid.UUID, sellerID uuid.UUID) error
	GetProductsBySellerID(ctx context.Context, sellerID uuid.UUID) ([]*domain.Product, error)
	CountProductsBySellerID(ctx context.Context, sellerID uuid.UUID) (int, error)
	FindProductByID(ctx context.Context, id uuid.UUID) (*domain.Product, error)
	GetAllProducts(ctx context.Context, category string) ([]*domain.Product, error)
	GetRecommendations(ctx context.Context, buyerID uuid.UUID) ([]*domain.Product, error)
	DeductStock(ctx context.Context, productID uuid.UUID, quantity int) error
}

type OrderRepository interface {
	GetOrdersBySellerID(ctx context.Context, sellerID uuid.UUID) ([]*domain.Order, error)
	CountOrdersBySellerID(ctx context.Context, sellerID uuid.UUID) (total, pending, completed int, err error)
	GetTotalRevenueBySellerID(ctx context.Context, sellerID uuid.UUID) (float64, error)
	GetOrdersByBuyerID(ctx context.Context, buyerID uuid.UUID) ([]*domain.Order, error)
	CreateOrder(ctx context.Context, order *domain.Order) (*domain.Order, error)
	UpdateOrderStatus(ctx context.Context, sellerID uuid.UUID, orderID uuid.UUID, status string) (*domain.Order, error)
}

type BuyerRepository interface {
	InsertCartItem(ctx context.Context, item *domain.CartItem) (*domain.CartItem, error)
	UpdateCartItemQuantity(ctx context.Context, itemID uuid.UUID, quantity int) (*domain.CartItem, error)
	GetCartItemsByBuyerID(ctx context.Context, buyerID uuid.UUID) ([]*domain.CartItem, error)
	FindCartItemByProductID(ctx context.Context, buyerID uuid.UUID, productID uuid.UUID) (*domain.CartItem, error)
	RemoveCartItem(ctx context.Context, itemID uuid.UUID, buyerID uuid.UUID) error
	ClearCart(ctx context.Context, buyerID uuid.UUID) error
	CreateBuyerTransaction(ctx context.Context, tx *domain.BuyerTransaction) (*domain.BuyerTransaction, error)
	GetBuyerTransactionsByBuyerID(ctx context.Context, buyerID uuid.UUID) ([]*domain.BuyerTransaction, error)
}
