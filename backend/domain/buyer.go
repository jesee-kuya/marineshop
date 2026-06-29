package domain

import (
	"time"

	"github.com/google/uuid"
)

type CartItem struct {
	ID        uuid.UUID `json:"id"`
	BuyerID   uuid.UUID `json:"buyer_id"`
	ProductID uuid.UUID `json:"product_id"`
	Product   *Product  `json:"product,omitempty"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BuyerTransaction struct {
	ID        uuid.UUID `json:"id"`
	BuyerID   uuid.UUID `json:"buyer_id"`
	OrderID   uuid.UUID `json:"order_id"`
	Amount    float64   `json:"amount"`
	Status    string    `json:"status"`
	Reference string    `json:"reference"`
	CreatedAt time.Time `json:"created_at"`
}
