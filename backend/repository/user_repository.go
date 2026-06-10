package repository

import (
	"context"

	"github.com/jesee-kuya/marineshop/domain"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	query := `
		INSERT INTO users (username, email, password, role)
		VALUES ($1, $2, $3, $4)
		RETURNING id, username, email, role
	`
	created := &domain.User{}
	err := r.db.QueryRowContext(ctx, query, user.Username, user.Email, user.Password, user.Role).
		Scan(&created.ID, &created.Username, &created.Email, &created.Role)
	if err != nil {
		return nil, err
	}
	return created, nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := `SELECT id, username, email, password, role FROM users WHERE email = $1`
	user := &domain.User{}
	err := r.db.QueryRowContext(ctx, query, email).
		Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return nil, err
	}
	return user, nil
}
