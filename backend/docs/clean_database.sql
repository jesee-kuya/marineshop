-- Run this as a superuser (e.g. postgres) to fully wipe and recreate the database.
-- Usage: psql -U postgres -f docs/clean_database.sql

-- 1. Drop all connections to the database
SELECT pg_terminate_backend(pid)
FROM pg_stat_activity
WHERE datname = 'marineshop' AND pid <> pg_backend_pid();

-- 2. Drop the database and user
DROP DATABASE IF EXISTS marineshop;
DROP USER IF EXISTS marineshop_user;

-- 3. Recreate user and database (same as init_database.sql)
CREATE USER marineshop_user WITH PASSWORD 'strong_password_here';
CREATE DATABASE marineshop;
ALTER DATABASE marineshop OWNER TO marineshop_user;
GRANT ALL PRIVILEGES ON DATABASE marineshop TO marineshop_user;
