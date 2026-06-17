package repository

import "github.com/jmoiron/sqlx"

type userRepository struct {
	db *sqlx.DB
}
