package model

import "time"

type User struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	Status    string `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
