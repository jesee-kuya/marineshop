package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jesee-kuya/marineshop/domain"
)

func (a *Admin) GetPendingKYCs(ctx context.Context) ([]*domain.SellerKYC, error) {
	return a.KYCRepo.GetPendingKYCs(ctx)
}

func (a *Admin) ApproveKYC(ctx context.Context, kycID uuid.UUID) (*domain.SellerKYC, error) {
	return a.KYCRepo.UpdateKYCStatus(ctx, kycID, "approved")
}

func (a *Admin) RejectKYC(ctx context.Context, kycID uuid.UUID) (*domain.SellerKYC, error) {
	return a.KYCRepo.UpdateKYCStatus(ctx, kycID, "rejected")
}
