package db

import (
	"database/sql"
	"os"
	"path/filepath"

	"github.com/jesee-kuya/marineshop/pkg/db/sqlite"
)

func Init() (*sql.DB, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	dbPath := filepath.Join(currentDir, "pkg", "db", "shop.db")
	migrationsPath := filepath.Join(currentDir, "pkg", "db", "sqlite")

	db, err := sqlite.InitDB(dbPath, migrationsPath)
	if err != nil {
		return nil, err
	}

	return db, nil
}
