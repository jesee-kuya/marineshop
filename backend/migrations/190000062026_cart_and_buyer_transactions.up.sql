CREATE TABLE IF NOT EXISTS cart_items (
    id         UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
    buyer_id   UUID        NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    product_id UUID        NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    quantity   INT         NOT NULL CHECK (quantity > 0),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (buyer_id, product_id)
);

CREATE TABLE IF NOT EXISTS buyer_transactions (
    id         UUID           PRIMARY KEY DEFAULT gen_random_uuid(),
    buyer_id   UUID           NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    order_id   UUID           NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
    amount     NUMERIC(19, 4) NOT NULL CHECK (amount > 0),
    status     VARCHAR(50)    NOT NULL DEFAULT 'pending',
    reference  TEXT           NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ    NOT NULL DEFAULT NOW()
);
