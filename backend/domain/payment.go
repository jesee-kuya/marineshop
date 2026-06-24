package domain

import (
	"time"

	"github.com/google/uuid"
)

type SellerPaymentAccount struct {
	ID            uuid.UUID `json:"id"`
	SellerID      uuid.UUID `json:"seller_id"`
	WalletType    string    `json:"wallet_type"`
	AccountName   string    `json:"account_name"`
	PhoneNumber   *string   `json:"phone_number,omitempty"`
	BankName      *string   `json:"bank_name,omitempty"`
	BankCode      *string   `json:"bank_code,omitempty"`
	AccountNumber *string   `json:"account_number,omitempty"`
	CryptoAddress *string   `json:"crypto_address,omitempty"`
	CryptoNetwork *string   `json:"crypto_network,omitempty"`
	IsDefault     bool      `json:"is_default"`
	IsActive      bool      `json:"is_active"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
