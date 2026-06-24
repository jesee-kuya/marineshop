package repository

import "github.com/jmoiron/sqlx"

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}
