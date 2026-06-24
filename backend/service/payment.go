package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jesee-kuya/marineshop/domain"
)

func (m *Money) SetUpPayment(ctx context.Context, userID uuid.UUID, req *domain.SetUpPaymentRequest) (*domain.SellerPaymentAccount, error) {
	sellerKYC, err := m.KYCRepo.FindKYCByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if sellerKYC == nil {
		return nil, domain.ErrSellerKYCNotFound
	}

	if err := ValidatePaymentDetails(req); err != nil {
		return nil, err
	}

	account := &domain.SellerPaymentAccount{
		SellerID:    userID,
		WalletType:  req.WalletType,
		AccountName: req.AccountName,
		IsDefault:   req.IsDefault,
	}

	switch req.WalletType {
	case "mpesa", "airtel_money":
		account.PhoneNumber = &req.PhoneNumber
	case "bank":
		account.BankName = &req.BankName
		account.BankCode = &req.BankCode
		account.AccountNumber = &req.AccountNumber
	case "crypto":
		account.CryptoAddress = &req.CryptoAddress
		account.CryptoNetwork = &req.CryptoNetwork
	}

	return m.MoneyRepo.CreateSellerPaymentAccount(ctx, account)
}

func ValidatePaymentDetails(req *domain.SetUpPaymentRequest) error {
	switch req.WalletType {
	case "mpesa", "airtel_money":
		if req.PhoneNumber == "" {
			return fmt.Errorf("%w: phone_number is required for %s", domain.ErrInvalidPaymentDetails, req.WalletType)
		}
	case "bank":
		if req.BankName == "" || req.BankCode == "" || req.AccountNumber == "" {
			return fmt.Errorf("%w: bank_name, bank_code, and account_number are required for bank payments", domain.ErrInvalidPaymentDetails)
		}
	case "crypto":
		if req.CryptoAddress == "" || req.CryptoNetwork == "" {
			return fmt.Errorf("%w: crypto_address and crypto_network are required for crypto payments", domain.ErrInvalidPaymentDetails)
		}
	}
	return nil
}
