package main

import (
	"database/sql"
	"fmt"
	"time"
)

type IAccounts interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	// UpdateAccount(*Account) error
	GetAccounts() ([]*Account, error)
	GetAccountByID(int) (*Account, error)
	// GetAccountByNumber(int) (*Account, error)
}

type Account struct {
	ID                int       `json:"id,omitempty"`
	FirstName         string    `json:"firstName"`
	LastName          string    `json:"lastName"`
	Number            int64     `json:"number"`
	EncryptedPassword string    `json:"-"`
	Balance           int64     `json:"balance"`
	CreatedAt         time.Time `json:"createdAt"`
}

type AccountStore struct {
	db *sql.DB
}

func (accStore *AccountStore) CreateAccount(account *Account) error {
	sqlQuery := `
	INSERT INTO accounts
		(first_name, last_name, number, encrypted_password, balance, created_at)
		VALUES(?, ?, ?, ?, ?, ?)
	`
	_, err := accStore.db.Query(
		sqlQuery,
		account.FirstName,
		account.LastName,
		account.Number,
		account.EncryptedPassword,
		account.Balance,
		account.CreatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (accStore *AccountStore) GetAccounts() ([]*Account, error) {
	sqlQuery := `SELECT * FROM accounts`
	rows, err := accStore.db.Query(sqlQuery)
	if err != nil {
		return nil, err
	}
	accounts := []*Account{}
	for rows.Next() {
		account, err := scanRowsIntoAccount(rows)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (accStore *AccountStore) GetAccountByID(id int) (*Account, error) {
	sqlQuery := `SELECT * FROM accounts WHERE id=?`
	rows, err := accStore.db.Query(sqlQuery, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanRowsIntoAccount(rows)
	}

	return nil, fmt.Errorf("Account %d not found", id)
}

func (accStore *AccountStore) DeleteAccount(id int) error {
	sqlQuery := `DELETE from accounts WHERE id=?`
	_, err := accStore.db.Query(sqlQuery, id)
	return err
}

func scanRowsIntoAccount(rows *sql.Rows) (*Account, error) {
	account := new(Account)
	var rawTimestamp []byte
	err := rows.Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.Number,
		&account.EncryptedPassword,
		&account.Balance,
		&rawTimestamp,
	)
	if err != nil {
		return nil, err
	}

	account.CreatedAt, err = time.Parse("2006-01-02 15:04:05", string(rawTimestamp))
	return account, err
}
