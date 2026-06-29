package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jesee-kuya/marineshop/domain"
)

func (r *moneyRepository) CreateSellerPaymentAccount(ctx context.Context, account *domain.SellerPaymentAccount) (*domain.SellerPaymentAccount, error) {
	query := `
		INSERT INTO seller_payment_accounts
			(seller_id, wallet_type, account_name, phone_number, bank_name, bank_code, account_number, crypto_address, crypto_network, is_default)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id, seller_id, wallet_type, account_name, phone_number, bank_name, bank_code, account_number, crypto_address, crypto_network, is_default, is_active, created_at, updated_at
	`
	var phoneNumber, bankName, bankCode, accountNumber, cryptoAddress, cryptoNetwork sql.NullString

	created := &domain.SellerPaymentAccount{}
	err := r.db.QueryRowContext(ctx, query,
		account.SellerID,
		account.WalletType,
		account.AccountName,
		toNullString(account.PhoneNumber),
		toNullString(account.BankName),
		toNullString(account.BankCode),
		toNullString(account.AccountNumber),
		toNullString(account.CryptoAddress),
		toNullString(account.CryptoNetwork),
		account.IsDefault,
	).Scan(
		&created.ID,
		&created.SellerID,
		&created.WalletType,
		&created.AccountName,
		&phoneNumber,
		&bankName,
		&bankCode,
		&accountNumber,
		&cryptoAddress,
		&cryptoNetwork,
		&created.IsDefault,
		&created.IsActive,
		&created.CreatedAt,
		&created.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	created.PhoneNumber = fromNullString(phoneNumber)
	created.BankName = fromNullString(bankName)
	created.BankCode = fromNullString(bankCode)
	created.AccountNumber = fromNullString(accountNumber)
	created.CryptoAddress = fromNullString(cryptoAddress)
	created.CryptoNetwork = fromNullString(cryptoNetwork)

	return created, nil
}

func (r *moneyRepository) FindPaymentAccountsBySellerID(ctx context.Context, sellerID uuid.UUID) ([]*domain.SellerPaymentAccount, error) {
	query := `
		SELECT id, seller_id, wallet_type, account_name, phone_number, bank_name, bank_code, account_number, crypto_address, crypto_network, is_default, is_active, created_at, updated_at
		FROM seller_payment_accounts
		WHERE seller_id = $1 AND is_active = TRUE
		ORDER BY is_default DESC, created_at ASC
	`
	rows, err := r.db.QueryContext(ctx, query, sellerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accounts []*domain.SellerPaymentAccount
	for rows.Next() {
		var phoneNumber, bankName, bankCode, accountNumber, cryptoAddress, cryptoNetwork sql.NullString
		a := &domain.SellerPaymentAccount{}
		if err := rows.Scan(
			&a.ID, &a.SellerID, &a.WalletType, &a.AccountName,
			&phoneNumber, &bankName, &bankCode, &accountNumber,
			&cryptoAddress, &cryptoNetwork,
			&a.IsDefault, &a.IsActive, &a.CreatedAt, &a.UpdatedAt,
		); err != nil {
			return nil, err
		}
		a.PhoneNumber = fromNullString(phoneNumber)
		a.BankName = fromNullString(bankName)
		a.BankCode = fromNullString(bankCode)
		a.AccountNumber = fromNullString(accountNumber)
		a.CryptoAddress = fromNullString(cryptoAddress)
		a.CryptoNetwork = fromNullString(cryptoNetwork)
		accounts = append(accounts, a)
	}

	return accounts, rows.Err()
}

func (r *moneyRepository) CreateTransaction(ctx context.Context, tx *domain.SellerTransaction) (*domain.SellerTransaction, error) {
	query := `
		INSERT INTO seller_transactions (seller_id, type, amount, status, reference, description, payment_account_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, seller_id, type, amount, status, reference, description, payment_account_id, created_at
	`
	created := &domain.SellerTransaction{}
	var paymentAccountID uuid.NullUUID
	err := r.db.QueryRowContext(ctx, query,
		tx.SellerID, tx.Type, tx.Amount, tx.Status, tx.Reference, tx.Description, toNullUUID(tx.PaymentAccountID),
	).Scan(
		&created.ID, &created.SellerID, &created.Type, &created.Amount,
		&created.Status, &created.Reference, &created.Description, &paymentAccountID, &created.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	created.PaymentAccountID = fromNullUUID(paymentAccountID)
	return created, nil
}

func (r *moneyRepository) GetTransactionsBySellerID(ctx context.Context, sellerID uuid.UUID) ([]*domain.SellerTransaction, error) {
	query := `
		SELECT id, seller_id, type, amount, status, reference, description, payment_account_id, created_at
		FROM seller_transactions WHERE seller_id = $1
		ORDER BY created_at DESC
	`
	rows, err := r.db.QueryContext(ctx, query, sellerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []*domain.SellerTransaction
	for rows.Next() {
		var paymentAccountID uuid.NullUUID
		t := &domain.SellerTransaction{}
		if err := rows.Scan(
			&t.ID, &t.SellerID, &t.Type, &t.Amount,
			&t.Status, &t.Reference, &t.Description, &paymentAccountID, &t.CreatedAt,
		); err != nil {
			return nil, err
		}
		t.PaymentAccountID = fromNullUUID(paymentAccountID)
		transactions = append(transactions, t)
	}
	return transactions, rows.Err()
}

func (r *moneyRepository) FindPaymentAccountByID(ctx context.Context, id uuid.UUID) (*domain.SellerPaymentAccount, error) {
	query := `
		SELECT id, seller_id, wallet_type, account_name, phone_number, bank_name, bank_code, account_number, crypto_address, crypto_network, is_default, is_active, created_at, updated_at
		FROM seller_payment_accounts WHERE id = $1
	`
	var phoneNumber, bankName, bankCode, accountNumber, cryptoAddress, cryptoNetwork sql.NullString
	a := &domain.SellerPaymentAccount{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&a.ID, &a.SellerID, &a.WalletType, &a.AccountName,
		&phoneNumber, &bankName, &bankCode, &accountNumber,
		&cryptoAddress, &cryptoNetwork,
		&a.IsDefault, &a.IsActive, &a.CreatedAt, &a.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	a.PhoneNumber = fromNullString(phoneNumber)
	a.BankName = fromNullString(bankName)
	a.BankCode = fromNullString(bankCode)
	a.AccountNumber = fromNullString(accountNumber)
	a.CryptoAddress = fromNullString(cryptoAddress)
	a.CryptoNetwork = fromNullString(cryptoNetwork)
	return a, nil
}

func (r *moneyRepository) GetSellerBalance(ctx context.Context, sellerID uuid.UUID) (float64, error) {
	query := `
		SELECT COALESCE(SUM(CASE WHEN type = 'credit' THEN amount WHEN type = 'withdrawal' THEN -amount ELSE 0 END), 0)
		FROM seller_transactions WHERE seller_id = $1
	`
	var balance float64
	err := r.db.QueryRowContext(ctx, query, sellerID).Scan(&balance)
	return balance, err
}

func toNullString(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{}
	}
	return sql.NullString{String: *s, Valid: true}
}

func fromNullString(ns sql.NullString) *string {
	if !ns.Valid {
		return nil
	}
	return &ns.String
}

func toNullUUID(u *uuid.UUID) uuid.NullUUID {
	if u == nil {
		return uuid.NullUUID{}
	}
	return uuid.NullUUID{UUID: *u, Valid: true}
}

func fromNullUUID(nu uuid.NullUUID) *uuid.UUID {
	if !nu.Valid {
		return nil
	}
	return &nu.UUID
}
