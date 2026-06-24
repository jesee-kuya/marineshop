package domain

import "errors"

var (
	ErrEmailInUse              = errors.New("email already in use")
	ErrInvalidCredentials      = errors.New("invalid email or password")
	ErrUserNotFound            = errors.New("user not found")
	ErrKYCAlreadyExists        = errors.New("KYC already submitted")
	ErrSellerKYCNotFound       = errors.New("seller KYC not found, please complete personal KYC first")
	ErrBusinessKYCAlreadyExists = errors.New("business KYC already submitted")
)
