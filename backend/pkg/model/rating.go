package model

import "time"

type Rating struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	AdID      string    `json:"ad_id"`
	Rating    int       `json:"rating"`
	CreatedAt time.Time `json:"created_at"`
}
