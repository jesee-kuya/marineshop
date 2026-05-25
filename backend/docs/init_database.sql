-- 1. Create a new user (role)
CREATE USER marineshop_user WITH PASSWORD 'strong_password_here';

-- 2. Create the database
CREATE DATABASE marineshop;

-- 3. Give ownership of database to the new user
ALTER DATABASE marineshop OWNER TO marineshop_user;

-- 4. Grant privileges
GRANT ALL PRIVILEGES ON DATABASE marineshop TO marineshop_user;