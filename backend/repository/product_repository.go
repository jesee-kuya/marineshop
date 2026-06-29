package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jesee-kuya/marineshop/domain"
)

func (r *productRepository) CreateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	query := `
		INSERT INTO products (seller_id, name, description, price, stock, category, image_url)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, seller_id, name, description, price, stock, category, image_url, is_active, created_at, updated_at
	`
	created := &domain.Product{}
	err := r.db.QueryRowContext(ctx, query,
		product.SellerID, product.Name, product.Description,
		product.Price, product.Stock, product.Category, product.ImageURL,
	).Scan(
		&created.ID, &created.SellerID, &created.Name, &created.Description,
		&created.Price, &created.Stock, &created.Category, &created.ImageURL,
		&created.IsActive, &created.CreatedAt, &created.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return created, nil
}

func (r *productRepository) UpdateProduct(ctx context.Context, id uuid.UUID, sellerID uuid.UUID, req *domain.UpdateProductRequest) (*domain.Product, error) {
	query := `
		UPDATE products
		SET name = $1, description = $2, price = $3, stock = $4, category = $5, image_url = $6, updated_at = NOW()
		WHERE id = $7 AND seller_id = $8 AND is_active = TRUE
		RETURNING id, seller_id, name, description, price, stock, category, image_url, is_active, created_at, updated_at
	`
	updated := &domain.Product{}
	err := r.db.QueryRowContext(ctx, query,
		req.Name, req.Description, req.Price, req.Stock, req.Category, req.ImageURL, id, sellerID,
	).Scan(
		&updated.ID, &updated.SellerID, &updated.Name, &updated.Description,
		&updated.Price, &updated.Stock, &updated.Category, &updated.ImageURL,
		&updated.IsActive, &updated.CreatedAt, &updated.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrProductNotFound
		}
		return nil, err
	}
	return updated, nil
}

func (r *productRepository) DeleteProduct(ctx context.Context, id uuid.UUID, sellerID uuid.UUID) error {
	query := `UPDATE products SET is_active = FALSE WHERE id = $1 AND seller_id = $2 AND is_active = TRUE`
	result, err := r.db.ExecContext(ctx, query, id, sellerID)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return domain.ErrProductNotFound
	}
	return nil
}

func (r *productRepository) GetAllProducts(ctx context.Context, category string) ([]*domain.Product, error) {
	query := `
		SELECT id, seller_id, name, description, price, stock, category, image_url, is_active, created_at, updated_at
		FROM products WHERE is_active = TRUE AND ($1 = '' OR category = $1)
		ORDER BY created_at DESC
	`
	rows, err := r.db.QueryContext(ctx, query, category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*domain.Product
	for rows.Next() {
		p := &domain.Product{}
		if err := rows.Scan(
			&p.ID, &p.SellerID, &p.Name, &p.Description,
			&p.Price, &p.Stock, &p.Category, &p.ImageURL,
			&p.IsActive, &p.CreatedAt, &p.UpdatedAt,
		); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, rows.Err()
}

func (r *productRepository) GetRecommendations(ctx context.Context, buyerID uuid.UUID) ([]*domain.Product, error) {
	query := `
		SELECT id, seller_id, name, description, price, stock, category, image_url, is_active, created_at, updated_at
		FROM products WHERE is_active = TRUE
		ORDER BY created_at DESC
		LIMIT 10
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*domain.Product
	for rows.Next() {
		p := &domain.Product{}
		if err := rows.Scan(
			&p.ID, &p.SellerID, &p.Name, &p.Description,
			&p.Price, &p.Stock, &p.Category, &p.ImageURL,
			&p.IsActive, &p.CreatedAt, &p.UpdatedAt,
		); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, rows.Err()
}

func (r *productRepository) DeductStock(ctx context.Context, productID uuid.UUID, quantity int) error {
	query := `
		UPDATE products SET stock = stock - $1, updated_at = NOW()
		WHERE id = $2 AND stock >= $1 AND is_active = TRUE
	`
	result, err := r.db.ExecContext(ctx, query, quantity, productID)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return domain.ErrOutOfStock
	}
	return nil
}

func (r *productRepository) CountProductsBySellerID(ctx context.Context, sellerID uuid.UUID) (int, error) {
	var count int
	err := r.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM products WHERE seller_id = $1 AND is_active = TRUE`, sellerID).Scan(&count)
	return count, err
}

func (r *productRepository) FindProductByID(ctx context.Context, id uuid.UUID) (*domain.Product, error) {
	query := `
		SELECT id, seller_id, name, description, price, stock, category, image_url, is_active, created_at, updated_at
		FROM products WHERE id = $1
	`
	p := &domain.Product{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&p.ID, &p.SellerID, &p.Name, &p.Description,
		&p.Price, &p.Stock, &p.Category, &p.ImageURL,
		&p.IsActive, &p.CreatedAt, &p.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return p, nil
}

func (r *productRepository) GetProductsBySellerID(ctx context.Context, sellerID uuid.UUID) ([]*domain.Product, error) {
	query := `
		SELECT id, seller_id, name, description, price, stock, category, image_url, is_active, created_at, updated_at
		FROM products WHERE seller_id = $1 AND is_active = TRUE
		ORDER BY created_at DESC
	`
	rows, err := r.db.QueryContext(ctx, query, sellerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*domain.Product
	for rows.Next() {
		p := &domain.Product{}
		if err := rows.Scan(
			&p.ID, &p.SellerID, &p.Name, &p.Description,
			&p.Price, &p.Stock, &p.Category, &p.ImageURL,
			&p.IsActive, &p.CreatedAt, &p.UpdatedAt,
		); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, rows.Err()
}
