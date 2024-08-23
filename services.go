package main

import (
	"log"
	"math/rand"
	"time"
)

func CreateAccountService(store *Storage, fiestName, lastName, password string) (*Account, error) {
	account := &Account{
		FirstName:         fiestName,
		LastName:          lastName,
		EncryptedPassword: password,
		Number:            int64(rand.Intn(1000000)),
		CreatedAt:         time.Now().UTC(),
	}
	if err := store.Accounts.CreateAccount(account); err != nil {
		log.Println("Failed to create account. Error", err)
		return nil, &NewError{Message: "Failed to create account"}
	}
	return account, nil
}

func GetAllAccountsService(store *Storage) ([]*Account, error) {
	accounts, err := store.Accounts.GetAccounts()
	if err != nil {
		log.Println("Failed to fectch accounts. Error", err)
		return nil, &NewError{Message: "Failed to fectch accounts"}
	}
	return accounts, nil
}

func GetAccountByIdService(store *Storage, id int) (*Account, error) {
	account, err := store.Accounts.GetAccountByID(id)
	if err != nil {
		log.Println("Failed to fectch account. Error", err)
		return nil, &NewError{Message: "Failed to fectch account"}
	}
	return account, nil
}

func DeleteAccountByIdService(store *Storage, id int) error {
	err := store.Accounts.DeleteAccount(id)
	if err != nil {
		log.Println("Failed to delete account. Error", err)
		return &NewError{Message: "Failed to delete account"}
	}
	return nil
}
