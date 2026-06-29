package repository

import (
	"context"
	"database/sql"
)

// DBTX is satisfied by both *sqlx.DB and *sqlx.Tx, enabling repositories to
// run inside or outside a database transaction without changing their API.
type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}

type userRepository struct {
	db DBTX
}

type kycRepository struct {
	db DBTX
}

type moneyRepository struct {
	db DBTX
}

type productRepository struct {
	db DBTX
}

type orderRepository struct {
	db DBTX
}

type buyerRepository struct {
	db DBTX
}
