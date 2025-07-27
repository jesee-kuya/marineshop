package model

import "time"

type Ad struct {
	ID            string    `json:"id"`
	SellerID      string    `json:"seller_id"`
	Description   string    `json:"description"`
	Price         float64   `json:"price"`
	ItemCount     int       `json:"item_count"`
	AverageRating float64   `json:"average_rating"`
	RatingsCount  int       `json:"ratings_count"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Media         []Media   `json:"media"`
}
