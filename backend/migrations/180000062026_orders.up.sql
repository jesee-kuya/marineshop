CREATE TYPE order_status AS ENUM ('pending', 'processing', 'shipped', 'completed', 'cancelled');

CREATE TABLE IF NOT EXISTS orders (
    id         UUID         PRIMARY KEY DEFAULT gen_random_uuid(),
    buyer_id   UUID         NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    seller_id  UUID         NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    product_id UUID         NOT NULL REFERENCES products(id),
    quantity   INT          NOT NULL CHECK (quantity > 0),
    total      NUMERIC(19, 4) NOT NULL CHECK (total > 0),
    status     order_status NOT NULL DEFAULT 'pending',
    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);
