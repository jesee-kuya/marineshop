include .env

DB_URL=$(DATABASE_URL)
MIGRATIONS_PATH=./migrations

up:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" up

down:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" down