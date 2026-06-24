package repository

import "github.com/jmoiron/sqlx"

type userRepository struct {
	db *sqlx.DB
}

type kycRepository struct {
	db *sqlx.DB
}

type moneyRepository struct {
	db *sqlx.DB
}
