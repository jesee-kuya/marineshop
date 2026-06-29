package domain

import "errors"

var (
	ErrEmailInUse               = errors.New("email already in use")
	ErrInvalidCredentials       = errors.New("invalid email or password")
	ErrUserNotFound             = errors.New("user not found")
	ErrKYCAlreadyExists         = errors.New("KYC already submitted")
	ErrSellerKYCNotFound        = errors.New("seller KYC not found, please complete personal KYC first")
	ErrBusinessKYCAlreadyExists = errors.New("business KYC already submitted")
	ErrInvalidPaymentDetails    = errors.New("invalid payment details")
	ErrProductNotFound          = errors.New("product not found")
	ErrInsufficientBalance      = errors.New("insufficient balance")
	ErrUnauthorized             = errors.New("unauthorized to perform this action")
	ErrCartItemNotFound         = errors.New("cart item not found")
	ErrOutOfStock               = errors.New("product is out of stock")
	ErrCartEmpty                = errors.New("cart is empty")
	ErrKYCNotFound              = errors.New("KYC not found")
	ErrOrderNotFound            = errors.New("order not found")
	ErrInvalidAdminSecret       = errors.New("invalid admin setup secret")
)
