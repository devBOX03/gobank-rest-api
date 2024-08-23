package main

import (
	"database/sql"
	"log/slog"

	_ "github.com/go-sql-driver/mysql"
)

func newDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/mybank")
	if err != nil {
		// slog.Error("Failed to connect to database", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	defer slog.Info("Successfully connected to MySQL DB.")
	return db, nil
}

type Storage struct {
	Accounts IAccounts
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Accounts: &AccountStore{db},
	}
}
