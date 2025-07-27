package model

type Cart struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	AdID      string `json:"ad_id"`
	Quantity  int    `json:"quantity"`
	CreatedAt string `json:"created_at"`
}
