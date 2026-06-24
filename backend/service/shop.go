package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jesee-kuya/marineshop/domain"
)

func (s *Seller) SetUpShop(ctx context.Context, userID uuid.UUID, req *domain.SetUpShopRequest) (*domain.BusinessKYC, error) {
	sellerKYC, err := s.KYCRepo.FindKYCByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if sellerKYC == nil {
		return nil, domain.ErrSellerKYCNotFound
	}

	existing, err := s.KYCRepo.FindBusinessKYCBySellerKYCID(ctx, sellerKYC.ID)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, domain.ErrBusinessKYCAlreadyExists
	}

	businessKYC := &domain.BusinessKYC{
		SellerKYCID:  sellerKYC.ID,
		BusinessName: req.BusinessName,
		DocumentType: req.DocumentType,
		Document:     req.Document,
	}

	return s.KYCRepo.CreateBusinessKYC(ctx, businessKYC)
}
