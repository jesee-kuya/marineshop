package domain

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID `json:"id"`
	SellerID    uuid.UUID `json:"seller_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	Category    string    `json:"category"`
	ImageURL    string    `json:"image_url"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Order struct {
	ID        uuid.UUID `json:"id"`
	BuyerID   uuid.UUID `json:"buyer_id"`
	SellerID  uuid.UUID `json:"seller_id"`
	ProductID uuid.UUID `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Total     float64   `json:"total"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SellerTransaction struct {
	ID               uuid.UUID  `json:"id"`
	SellerID         uuid.UUID  `json:"seller_id"`
	Type             string     `json:"type"`
	Amount           float64    `json:"amount"`
	Status           string     `json:"status"`
	Reference        string     `json:"reference"`
	Description      string     `json:"description"`
	PaymentAccountID *uuid.UUID `json:"payment_account_id,omitempty"`
	CreatedAt        time.Time  `json:"created_at"`
}

type SellerAnalytics struct {
	TotalProducts   int     `json:"total_products"`
	TotalOrders     int     `json:"total_orders"`
	TotalRevenue    float64 `json:"total_revenue"`
	PendingOrders   int     `json:"pending_orders"`
	CompletedOrders int     `json:"completed_orders"`
}

type SellerProfile struct {
	KYC         *SellerKYC   `json:"kyc"`
	BusinessKYC *BusinessKYC `json:"business_kyc"`
}
