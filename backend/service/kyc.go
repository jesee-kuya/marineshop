package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jesee-kuya/marineshop/domain"
)

func (s *Seller) CollectKYC(ctx context.Context, userID uuid.UUID, req *domain.CollectKYCRequest) (*domain.SellerKYC, error) {
	existing, err := s.KYCRepo.FindKYCByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, domain.ErrKYCAlreadyExists
	}

	kyc := &domain.SellerKYC{
		UserID:             userID,
		FullName:           req.FullName,
		PhoneNumber:        req.PhoneNumber,
		NationalID:         req.NationalID,
		NationalIDDocument: req.NationalIDDocument,
		Selfie:             req.Selfie,
		Location:           req.Location,
	}

	return s.KYCRepo.CreateSellerKYC(ctx, kyc)
}
