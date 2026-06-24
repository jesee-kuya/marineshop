package domain

import (
	"time"

	"github.com/google/uuid"
)

type SellerKYC struct {
	ID                 uuid.UUID `json:"id"`
	UserID             uuid.UUID `json:"user_id"`
	FullName           string    `json:"full_name"`
	PhoneNumber        string    `json:"phone_number"`
	NationalID         string    `json:"national_id"`
	NationalIDDocument string    `json:"national_id_document"`
	Selfie             string    `json:"selfie"`
	Location           string    `json:"location"`
	Status             string    `json:"status"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}
