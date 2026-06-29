package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jesee-kuya/marineshop/domain"
)

func (r *orderRepository) GetOrdersBySellerID(ctx context.Context, sellerID uuid.UUID) ([]*domain.Order, error) {
	query := `
		SELECT id, buyer_id, seller_id, product_id, quantity, total, status, created_at, updated_at
		FROM orders WHERE seller_id = $1
		ORDER BY created_at DESC
	`
	rows, err := r.db.QueryContext(ctx, query, sellerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*domain.Order
	for rows.Next() {
		o := &domain.Order{}
		if err := rows.Scan(
			&o.ID, &o.BuyerID, &o.SellerID, &o.ProductID,
			&o.Quantity, &o.Total, &o.Status, &o.CreatedAt, &o.UpdatedAt,
		); err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}
	return orders, rows.Err()
}

func (r *orderRepository) CountOrdersBySellerID(ctx context.Context, sellerID uuid.UUID) (total, pending, completed int, err error) {
	query := `
		SELECT
			COUNT(*) AS total,
			COUNT(*) FILTER (WHERE status = 'pending') AS pending,
			COUNT(*) FILTER (WHERE status = 'completed') AS completed
		FROM orders WHERE seller_id = $1
	`
	err = r.db.QueryRowContext(ctx, query, sellerID).Scan(&total, &pending, &completed)
	return
}

func (r *orderRepository) GetTotalRevenueBySellerID(ctx context.Context, sellerID uuid.UUID) (float64, error) {
	query := `SELECT COALESCE(SUM(total), 0) FROM orders WHERE seller_id = $1 AND status = 'completed'`
	var revenue float64
	err := r.db.QueryRowContext(ctx, query, sellerID).Scan(&revenue)
	return revenue, err
}

func (r *orderRepository) GetOrdersByBuyerID(ctx context.Context, buyerID uuid.UUID) ([]*domain.Order, error) {
	query := `
		SELECT id, buyer_id, seller_id, product_id, quantity, total, status, created_at, updated_at
		FROM orders WHERE buyer_id = $1
		ORDER BY created_at DESC
	`
	rows, err := r.db.QueryContext(ctx, query, buyerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*domain.Order
	for rows.Next() {
		o := &domain.Order{}
		if err := rows.Scan(
			&o.ID, &o.BuyerID, &o.SellerID, &o.ProductID,
			&o.Quantity, &o.Total, &o.Status, &o.CreatedAt, &o.UpdatedAt,
		); err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}
	return orders, rows.Err()
}

func (r *orderRepository) UpdateOrderStatus(ctx context.Context, sellerID uuid.UUID, orderID uuid.UUID, status string) (*domain.Order, error) {
	query := `
		UPDATE orders SET status = $1, updated_at = NOW()
		WHERE id = $2 AND seller_id = $3
		RETURNING id, buyer_id, seller_id, product_id, quantity, total, status, created_at, updated_at
	`
	o := &domain.Order{}
	err := r.db.QueryRowContext(ctx, query, status, orderID, sellerID).Scan(
		&o.ID, &o.BuyerID, &o.SellerID, &o.ProductID,
		&o.Quantity, &o.Total, &o.Status, &o.CreatedAt, &o.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrOrderNotFound
		}
		return nil, err
	}
	return o, nil
}

func (r *orderRepository) CreateOrder(ctx context.Context, order *domain.Order) (*domain.Order, error) {
	query := `
		INSERT INTO orders (buyer_id, seller_id, product_id, quantity, total, status)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, buyer_id, seller_id, product_id, quantity, total, status, created_at, updated_at
	`
	created := &domain.Order{}
	err := r.db.QueryRowContext(ctx, query,
		order.BuyerID, order.SellerID, order.ProductID,
		order.Quantity, order.Total, order.Status,
	).Scan(
		&created.ID, &created.BuyerID, &created.SellerID, &created.ProductID,
		&created.Quantity, &created.Total, &created.Status,
		&created.CreatedAt, &created.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return created, nil
}
