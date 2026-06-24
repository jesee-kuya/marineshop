package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jesee-kuya/marineshop/domain"
)

func (r *kycRepository) CreateSellerKYC(ctx context.Context, kyc *domain.SellerKYC) (*domain.SellerKYC, error) {
	query := `
		INSERT INTO seller_kyc (user_id, full_name, phone_number, national_id, national_id_document, selfie, location)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, user_id, full_name, phone_number, national_id, national_id_document, selfie, location, status, created_at, updated_at
	`
	created := &domain.SellerKYC{}
	err := r.db.QueryRowContext(ctx, query,
		kyc.UserID, kyc.FullName, kyc.PhoneNumber, kyc.NationalID,
		kyc.NationalIDDocument, kyc.Selfie, kyc.Location,
	).Scan(
		&created.ID, &created.UserID, &created.FullName, &created.PhoneNumber,
		&created.NationalID, &created.NationalIDDocument, &created.Selfie,
		&created.Location, &created.Status, &created.CreatedAt, &created.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return created, nil
}

func (r *kycRepository) FindKYCByUserID(ctx context.Context, userID uuid.UUID) (*domain.SellerKYC, error) {
	query := `
		SELECT id, user_id, full_name, phone_number, national_id, national_id_document, selfie, location, status, created_at, updated_at
		FROM seller_kyc WHERE user_id = $1
	`
	kyc := &domain.SellerKYC{}
	err := r.db.QueryRowContext(ctx, query, userID).Scan(
		&kyc.ID, &kyc.UserID, &kyc.FullName, &kyc.PhoneNumber,
		&kyc.NationalID, &kyc.NationalIDDocument, &kyc.Selfie,
		&kyc.Location, &kyc.Status, &kyc.CreatedAt, &kyc.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return kyc, nil
}
