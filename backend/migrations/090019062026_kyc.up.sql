CREATE TYPE kyc_status AS ENUM ('pending', 'approved', 'rejected');

CREATE TYPE business_document_type AS ENUM ('permit', 'certificate', 'incorporation_letter');

CREATE TABLE IF NOT EXISTS seller_kyc (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL UNIQUE REFERENCES users(id) ON DELETE CASCADE,

    full_name VARCHAR(255) NOT NULL,
    phone_number VARCHAR(20) NOT NULL,
    national_id VARCHAR(50) NOT NULL UNIQUE,
    national_id_document TEXT NOT NULL,
    selfie TEXT NOT NULL,
    location TEXT NOT NULL,

    status kyc_status NOT NULL DEFAULT 'pending',

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS business_kyc (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    seller_kyc_id UUID NOT NULL UNIQUE REFERENCES seller_kyc(id) ON DELETE CASCADE,

    business_name VARCHAR(255) NOT NULL,
    document_type business_document_type NOT NULL,
    document TEXT NOT NULL,

    status kyc_status NOT NULL DEFAULT 'pending',

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
