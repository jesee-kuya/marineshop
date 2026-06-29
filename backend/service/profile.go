package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jesee-kuya/marineshop/domain"
)

func (s *Seller) GetProfile(ctx context.Context, userID uuid.UUID) (*domain.SellerProfile, error) {
	kyc, err := s.KYCRepo.FindKYCByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if kyc == nil {
		return nil, domain.ErrSellerKYCNotFound
	}

	businessKYC, err := s.KYCRepo.FindBusinessKYCBySellerKYCID(ctx, kyc.ID)
	if err != nil {
		return nil, err
	}

	return &domain.SellerProfile{
		KYC:         kyc,
		BusinessKYC: businessKYC,
	}, nil
}
