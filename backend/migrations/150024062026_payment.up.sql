CREATE TYPE wallet_type AS ENUM ('mpesa', 'airtel_money', 'bank', 'crypto');

CREATE TYPE transaction_direction AS ENUM ('inbound', 'outbound');

CREATE TYPE transaction_status AS ENUM ('pending', 'processing', 'completed', 'failed', 'reversed');

CREATE TYPE escrow_status AS ENUM ('held', 'released', 'refunded', 'disputed');

-- Seller's registered payout destinations.
-- SetUpPayment writes here. A seller can register one account per wallet type,
-- or multiple accounts, with is_default marking the preferred one.
CREATE TABLE IF NOT EXISTS seller_payment_accounts (
    id           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    seller_id    UUID        NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    wallet_type  wallet_type NOT NULL,

    -- Display name on the account (e.g. registered M-Pesa name, bank account holder)
    account_name VARCHAR(255) NOT NULL,

    -- Mobile money (mpesa, airtel_money)
    phone_number VARCHAR(20),

    -- Bank
    bank_name      VARCHAR(255),
    bank_code      VARCHAR(50),
    account_number VARCHAR(100),

    -- Crypto
    crypto_address VARCHAR(255),
    crypto_network VARCHAR(100),

    is_default BOOLEAN     NOT NULL DEFAULT FALSE,
    is_active  BOOLEAN     NOT NULL DEFAULT TRUE,

    -- Enforce required fields per wallet type at the DB level
    CONSTRAINT chk_spa_mobile_money CHECK (
        wallet_type NOT IN ('mpesa', 'airtel_money') OR phone_number IS NOT NULL
    ),
    CONSTRAINT chk_spa_bank CHECK (
        wallet_type != 'bank'
        OR (bank_name IS NOT NULL AND bank_code IS NOT NULL AND account_number IS NOT NULL)
    ),
    CONSTRAINT chk_spa_crypto CHECK (
        wallet_type != 'crypto'
        OR (crypto_address IS NOT NULL AND crypto_network IS NOT NULL)
    ),

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- The platform's own accounts used to receive buyer payments and hold escrow,
-- then disburse to sellers. Seeded and managed by admin — not exposed to sellers.
CREATE TABLE IF NOT EXISTS platform_accounts (
    id          UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
    wallet_type wallet_type NOT NULL,

    account_name   VARCHAR(255) NOT NULL,
    phone_number   VARCHAR(20),
    bank_name      VARCHAR(255),
    bank_code      VARCHAR(50),
    account_number VARCHAR(100),
    crypto_address VARCHAR(255),
    crypto_network VARCHAR(100),

    -- Human-readable label, e.g. "Primary KES M-Pesa escrow"
    description TEXT,
    is_active   BOOLEAN NOT NULL DEFAULT TRUE,

    CONSTRAINT chk_pa_mobile_money CHECK (
        wallet_type NOT IN ('mpesa', 'airtel_money') OR phone_number IS NOT NULL
    ),
    CONSTRAINT chk_pa_bank CHECK (
        wallet_type != 'bank'
        OR (bank_name IS NOT NULL AND bank_code IS NOT NULL AND account_number IS NOT NULL)
    ),
    CONSTRAINT chk_pa_crypto CHECK (
        wallet_type != 'crypto'
        OR (crypto_address IS NOT NULL AND crypto_network IS NOT NULL)
    ),

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Every money movement, inbound (buyer → platform) or outbound (platform → seller).
-- external_reference is the provider's confirmation code (M-Pesa receipt, bank ref, tx hash).
-- metadata holds the raw provider callback payload for audit and reconciliation.
CREATE TABLE IF NOT EXISTS transactions (
    id                 UUID               PRIMARY KEY DEFAULT gen_random_uuid(),
    external_reference VARCHAR(255)       UNIQUE,
    amount             NUMERIC(19, 4)     NOT NULL CHECK (amount > 0),
    currency           VARCHAR(10)        NOT NULL DEFAULT 'KES',
    direction          transaction_direction NOT NULL,
    status             transaction_status NOT NULL DEFAULT 'pending',
    wallet_type        wallet_type        NOT NULL,
    sender_account     VARCHAR(255)       NOT NULL,
    receiver_account   VARCHAR(255)       NOT NULL,
    description        TEXT,
    metadata           JSONB,
    created_at         TIMESTAMPTZ        NOT NULL DEFAULT NOW(),
    updated_at         TIMESTAMPTZ        NOT NULL DEFAULT NOW()
);

-- Ties buyer payment to a deferred seller disbursement.
-- order_id will FK to orders(id) once that table is introduced.
-- inbound_transaction_id  = the record of the buyer paying the platform.
-- outbound_transaction_id = the record of the platform paying the seller (set on release).
CREATE TABLE IF NOT EXISTS escrow (
    id                      UUID          PRIMARY KEY DEFAULT gen_random_uuid(),
    order_id                UUID          NOT NULL UNIQUE,
    buyer_id                UUID          NOT NULL REFERENCES users(id),
    seller_id               UUID          NOT NULL REFERENCES users(id),
    platform_account_id     UUID          NOT NULL REFERENCES platform_accounts(id),
    amount                  NUMERIC(19,4) NOT NULL CHECK (amount > 0),
    currency                VARCHAR(10)   NOT NULL DEFAULT 'KES',
    status                  escrow_status NOT NULL DEFAULT 'held',
    inbound_transaction_id  UUID          REFERENCES transactions(id),
    outbound_transaction_id UUID          REFERENCES transactions(id),
    held_at                 TIMESTAMPTZ,
    released_at             TIMESTAMPTZ,
    created_at              TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    updated_at              TIMESTAMPTZ   NOT NULL DEFAULT NOW()
);
