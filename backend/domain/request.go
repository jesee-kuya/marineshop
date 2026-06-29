package domain

import "github.com/google/uuid"

type SignupRequest struct {
	Username    string `json:"username" binding:"required,min=3,max=255"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=8"`
	Role        string `json:"role" binding:"required,oneof=buyer seller admin"`
	AdminSecret string `json:"admin_secret"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=8"`
}

type ResetPasswordRequest struct {
	Email           string `json:"email" binding:"required,email"`
	CurrentPassword string `json:"current_password" binding:"required"`
	NewPassword     string `json:"new_password" binding:"required,min=8"`
}

type CollectKYCRequest struct {
	FullName           string `json:"full_name" binding:"required"`
	PhoneNumber        string `json:"phone_number" binding:"required"`
	NationalID         string `json:"national_id" binding:"required"`
	NationalIDDocument string `json:"national_id_document" binding:"required"`
	Selfie             string `json:"selfie" binding:"required"`
	Location           string `json:"location" binding:"required"`
}

type SetUpShopRequest struct {
	BusinessName string `json:"business_name" binding:"required"`
	DocumentType string `json:"document_type" binding:"required,oneof=permit certificate incorporation_letter"`
	Document     string `json:"document" binding:"required"`
}

type SetUpPaymentRequest struct {
	WalletType    string `json:"wallet_type" binding:"required,oneof=mpesa airtel_money bank crypto"`
	AccountName   string `json:"account_name" binding:"required"`
	PhoneNumber   string `json:"phone_number"`
	BankName      string `json:"bank_name"`
	BankCode      string `json:"bank_code"`
	AccountNumber string `json:"account_number"`
	CryptoAddress string `json:"crypto_address"`
	CryptoNetwork string `json:"crypto_network"`
	IsDefault     bool   `json:"is_default"`
}

type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Stock       int     `json:"stock" binding:"required,gte=0"`
	Category    string  `json:"category" binding:"required"`
	ImageURL    string  `json:"image_url"`
}

type UpdateProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Stock       int     `json:"stock" binding:"required,gte=0"`
	Category    string  `json:"category" binding:"required"`
	ImageURL    string  `json:"image_url"`
}

type WithdrawRequest struct {
	Amount           float64   `json:"amount" binding:"required,gt=0"`
	PaymentAccountID uuid.UUID `json:"payment_account_id" binding:"required"`
}

type AddCartItemRequest struct {
	ProductID uuid.UUID `json:"product_id" binding:"required"`
	Quantity  int       `json:"quantity" binding:"required,gt=0"`
}

type CheckoutRequest struct {
	PaymentMethod string `json:"payment_method" binding:"required,oneof=mpesa airtel_money bank crypto"`
}

type UpdateOrderStatusRequest struct {
	Status string `json:"status" binding:"required,oneof=processing shipped completed cancelled"`
}
