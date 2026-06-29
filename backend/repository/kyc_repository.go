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

func (r *kycRepository) CreateBusinessKYC(ctx context.Context, kyc *domain.BusinessKYC) (*domain.BusinessKYC, error) {
	query := `
		INSERT INTO business_kyc (seller_kyc_id, business_name, document_type, document)
		VALUES ($1, $2, $3, $4)
		RETURNING id, seller_kyc_id, business_name, document_type, document, status, created_at, updated_at
	`
	created := &domain.BusinessKYC{}
	err := r.db.QueryRowContext(ctx, query,
		kyc.SellerKYCID, kyc.BusinessName, kyc.DocumentType, kyc.Document,
	).Scan(
		&created.ID, &created.SellerKYCID, &created.BusinessName,
		&created.DocumentType, &created.Document, &created.Status,
		&created.CreatedAt, &created.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return created, nil
}

func (r *kycRepository) GetPendingKYCs(ctx context.Context) ([]*domain.SellerKYC, error) {
	query := `
		SELECT id, user_id, full_name, phone_number, national_id, national_id_document, selfie, location, status, created_at, updated_at
		FROM seller_kyc WHERE status = 'pending'
		ORDER BY created_at ASC
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var kycs []*domain.SellerKYC
	for rows.Next() {
		kyc := &domain.SellerKYC{}
		if err := rows.Scan(
			&kyc.ID, &kyc.UserID, &kyc.FullName, &kyc.PhoneNumber,
			&kyc.NationalID, &kyc.NationalIDDocument, &kyc.Selfie,
			&kyc.Location, &kyc.Status, &kyc.CreatedAt, &kyc.UpdatedAt,
		); err != nil {
			return nil, err
		}
		kycs = append(kycs, kyc)
	}
	return kycs, rows.Err()
}

func (r *kycRepository) UpdateKYCStatus(ctx context.Context, id uuid.UUID, status string) (*domain.SellerKYC, error) {
	query := `
		UPDATE seller_kyc SET status = $1, updated_at = NOW()
		WHERE id = $2
		RETURNING id, user_id, full_name, phone_number, national_id, national_id_document, selfie, location, status, created_at, updated_at
	`
	updated := &domain.SellerKYC{}
	err := r.db.QueryRowContext(ctx, query, status, id).Scan(
		&updated.ID, &updated.UserID, &updated.FullName, &updated.PhoneNumber,
		&updated.NationalID, &updated.NationalIDDocument, &updated.Selfie,
		&updated.Location, &updated.Status, &updated.CreatedAt, &updated.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrKYCNotFound
		}
		return nil, err
	}
	return updated, nil
}

func (r *kycRepository) FindBusinessKYCBySellerKYCID(ctx context.Context, sellerKYCID uuid.UUID) (*domain.BusinessKYC, error) {
	query := `
		SELECT id, seller_kyc_id, business_name, document_type, document, status, created_at, updated_at
		FROM business_kyc WHERE seller_kyc_id = $1
	`
	kyc := &domain.BusinessKYC{}
	err := r.db.QueryRowContext(ctx, query, sellerKYCID).Scan(
		&kyc.ID, &kyc.SellerKYCID, &kyc.BusinessName,
		&kyc.DocumentType, &kyc.Document, &kyc.Status,
		&kyc.CreatedAt, &kyc.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return kyc, nil
}
