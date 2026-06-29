CREATE TABLE IF NOT EXISTS seller_transactions (
    id                 UUID           PRIMARY KEY DEFAULT gen_random_uuid(),
    seller_id          UUID           NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    type               VARCHAR(50)    NOT NULL CHECK (type IN ('credit', 'withdrawal')),
    amount             NUMERIC(19, 4) NOT NULL CHECK (amount > 0),
    status             VARCHAR(50)    NOT NULL DEFAULT 'pending',
    reference          TEXT           NOT NULL DEFAULT '',
    description        TEXT           NOT NULL DEFAULT '',
    payment_account_id UUID           REFERENCES seller_payment_accounts(id),
    created_at         TIMESTAMPTZ    NOT NULL DEFAULT NOW()
);
