package model

import "time"

type Review struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	AdID      string    `json:"ad_id"`
	Review    string    `json:"review"`
	CreatedAt time.Time `json:"created_at"`
}
