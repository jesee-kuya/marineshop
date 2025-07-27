CREATE TABLE IF NOT EXISTS users (
    id TEXT PRIMARY KEY NOT NULL UNIQUE,
    email TEXT NOT NULL UNIQUE,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    role TEXT CHECK(role IN ('admin', 'user', 'seller')) DEFAULT 'user',
    status TEXT CHECK(status IN ('active', 'inactive')) DEFAULT 'active',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
)