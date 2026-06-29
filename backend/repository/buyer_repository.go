package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jesee-kuya/marineshop/domain"
)

func (r *buyerRepository) InsertCartItem(ctx context.Context, item *domain.CartItem) (*domain.CartItem, error) {
	query := `
		INSERT INTO cart_items (buyer_id, product_id, quantity)
		VALUES ($1, $2, $3)
		RETURNING id, buyer_id, product_id, quantity, created_at, updated_at
	`
	created := &domain.CartItem{}
	err := r.db.QueryRowContext(ctx, query, item.BuyerID, item.ProductID, item.Quantity).Scan(
		&created.ID, &created.BuyerID, &created.ProductID, &created.Quantity,
		&created.CreatedAt, &created.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return created, nil
}

func (r *buyerRepository) UpdateCartItemQuantity(ctx context.Context, itemID uuid.UUID, quantity int) (*domain.CartItem, error) {
	query := `
		UPDATE cart_items SET quantity = $1, updated_at = NOW()
		WHERE id = $2
		RETURNING id, buyer_id, product_id, quantity, created_at, updated_at
	`
	updated := &domain.CartItem{}
	err := r.db.QueryRowContext(ctx, query, quantity, itemID).Scan(
		&updated.ID, &updated.BuyerID, &updated.ProductID, &updated.Quantity,
		&updated.CreatedAt, &updated.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return updated, nil
}

func (r *buyerRepository) GetCartItemsByBuyerID(ctx context.Context, buyerID uuid.UUID) ([]*domain.CartItem, error) {
	query := `
		SELECT
			c.id, c.buyer_id, c.product_id, c.quantity, c.created_at, c.updated_at,
			p.id, p.seller_id, p.name, p.description, p.price, p.stock, p.category, p.image_url, p.is_active, p.created_at, p.updated_at
		FROM cart_items c
		JOIN products p ON p.id = c.product_id
		WHERE c.buyer_id = $1
		ORDER BY c.created_at ASC
	`
	rows, err := r.db.QueryContext(ctx, query, buyerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*domain.CartItem
	for rows.Next() {
		item := &domain.CartItem{}
		p := &domain.Product{}
		if err := rows.Scan(
			&item.ID, &item.BuyerID, &item.ProductID, &item.Quantity, &item.CreatedAt, &item.UpdatedAt,
			&p.ID, &p.SellerID, &p.Name, &p.Description, &p.Price, &p.Stock, &p.Category, &p.ImageURL, &p.IsActive, &p.CreatedAt, &p.UpdatedAt,
		); err != nil {
			return nil, err
		}
		item.Product = p
		items = append(items, item)
	}
	return items, rows.Err()
}

func (r *buyerRepository) FindCartItemByProductID(ctx context.Context, buyerID uuid.UUID, productID uuid.UUID) (*domain.CartItem, error) {
	query := `
		SELECT id, buyer_id, product_id, quantity, created_at, updated_at
		FROM cart_items WHERE buyer_id = $1 AND product_id = $2
	`
	item := &domain.CartItem{}
	err := r.db.QueryRowContext(ctx, query, buyerID, productID).Scan(
		&item.ID, &item.BuyerID, &item.ProductID, &item.Quantity,
		&item.CreatedAt, &item.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return item, nil
}

func (r *buyerRepository) RemoveCartItem(ctx context.Context, itemID uuid.UUID, buyerID uuid.UUID) error {
	query := `DELETE FROM cart_items WHERE id = $1 AND buyer_id = $2`
	result, err := r.db.ExecContext(ctx, query, itemID, buyerID)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return domain.ErrCartItemNotFound
	}
	return nil
}

func (r *buyerRepository) ClearCart(ctx context.Context, buyerID uuid.UUID) error {
	query := `DELETE FROM cart_items WHERE buyer_id = $1`
	_, err := r.db.ExecContext(ctx, query, buyerID)
	return err
}

func (r *buyerRepository) CreateBuyerTransaction(ctx context.Context, tx *domain.BuyerTransaction) (*domain.BuyerTransaction, error) {
	query := `
		INSERT INTO buyer_transactions (buyer_id, order_id, amount, status, reference)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, buyer_id, order_id, amount, status, reference, created_at
	`
	created := &domain.BuyerTransaction{}
	err := r.db.QueryRowContext(ctx, query,
		tx.BuyerID, tx.OrderID, tx.Amount, tx.Status, tx.Reference,
	).Scan(
		&created.ID, &created.BuyerID, &created.OrderID,
		&created.Amount, &created.Status, &created.Reference, &created.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return created, nil
}

func (r *buyerRepository) GetBuyerTransactionsByBuyerID(ctx context.Context, buyerID uuid.UUID) ([]*domain.BuyerTransaction, error) {
	query := `
		SELECT id, buyer_id, order_id, amount, status, reference, created_at
		FROM buyer_transactions WHERE buyer_id = $1
		ORDER BY created_at DESC
	`
	rows, err := r.db.QueryContext(ctx, query, buyerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []*domain.BuyerTransaction
	for rows.Next() {
		t := &domain.BuyerTransaction{}
		if err := rows.Scan(
			&t.ID, &t.BuyerID, &t.OrderID,
			&t.Amount, &t.Status, &t.Reference, &t.CreatedAt,
		); err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}
	return transactions, rows.Err()
}
