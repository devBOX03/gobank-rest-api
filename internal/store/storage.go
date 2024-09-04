package store

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Storage struct {
	Accounts IAccounts
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Accounts: &AccountStore{db},
	}
}
