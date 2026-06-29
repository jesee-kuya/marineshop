package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
)

func NewUserRepository(db DBTX) UserRepository {
	return &userRepository{db: db}
}

func NewKYCRepository(db DBTX) KYCRepository {
	return &kycRepository{db: db}
}

func NewMoneyRepository(db DBTX) MoneyRepository {
	return &moneyRepository{db: db}
}

func NewProductRepository(db DBTX) ProductRepository {
	return &productRepository{db: db}
}

func NewOrderRepository(db DBTX) OrderRepository {
	return &orderRepository{db: db}
}

func NewBuyerRepository(db DBTX) BuyerRepository {
	return &buyerRepository{db: db}
}

type store struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) CheckoutStore {
	return &store{db: db}
}

func (s *store) ExecCheckoutTx(ctx context.Context, fn func(
	buyerRepo BuyerRepository,
	productRepo ProductRepository,
	orderRepo OrderRepository,
	moneyRepo MoneyRepository,
) error) error {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := fn(
		NewBuyerRepository(tx),
		NewProductRepository(tx),
		NewOrderRepository(tx),
		NewMoneyRepository(tx),
	); err != nil {
		return err
	}

	return tx.Commit()
}
