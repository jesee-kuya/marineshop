package repository

import "github.com/jmoiron/sqlx"

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}

func NewKYCRepository(db *sqlx.DB) KYCRepository {
	return &kycRepository{db: db}
}

func NewMoneyRepository(db *sqlx.DB) MoneyRepository {
	return &moneyRepository{db: db}
}
