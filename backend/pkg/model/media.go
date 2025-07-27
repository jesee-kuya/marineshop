package model

import "time"

type Media struct {
	ID        string    `json:"id"`
	ParentID  string    `json:"parent_id"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
