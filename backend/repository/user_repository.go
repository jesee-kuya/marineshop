package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jesee-kuya/marineshop/domain"
)

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
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (r *userRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	query := `SELECT id, username, email, password, role FROM users WHERE id = $1`
	user := &domain.User{}
	err := r.db.QueryRowContext(ctx, query, id).
		Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (r *userRepository) UpdatePassword(ctx context.Context, id uuid.UUID, hashedPassword string) error {
	query := `UPDATE users SET password = $1 WHERE id = $2`
	_, err := r.db.ExecContext(ctx, query, hashedPassword, id)
	return err
}
