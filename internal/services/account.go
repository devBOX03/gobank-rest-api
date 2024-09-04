package service

import (
	"log"
	"math/rand"
	"time"

	store "github.com/devBOX03/gobank-rest-api/internal/store"
	"github.com/devBOX03/gobank-rest-api/types"
)

func CreateAccountService(s *store.Storage, fiestName, lastName, password string) (*store.Account, error) {
	account := &store.Account{
		FirstName:         fiestName,
		LastName:          lastName,
		EncryptedPassword: password,
		Number:            int64(rand.Intn(1000000)),
		CreatedAt:         time.Now().UTC(),
	}
	if err := s.Accounts.CreateAccount(account); err != nil {
		log.Println("Failed to create account. Error", err)
		return nil, &types.NewError{Message: "Failed to create account"}
	}
	return account, nil
}

func GetAllAccountsService(s *store.Storage) ([]*store.Account, error) {
	accounts, err := s.Accounts.GetAccounts()
	if err != nil {
		log.Println("Failed to fectch accounts. Error", err)
		return nil, &types.NewError{Message: "Failed to fectch accounts"}
	}
	return accounts, nil
}

func GetAccountByIdService(s *store.Storage, id int) (*store.Account, error) {
	account, err := s.Accounts.GetAccountByID(id)
	if err != nil {
		log.Println("Failed to fectch account. Error", err)
		return nil, &types.NewError{Message: "Failed to fectch account"}
	}
	return account, nil
}

func DeleteAccountByIdService(s *store.Storage, id int) error {
	err := s.Accounts.DeleteAccount(id)
	if err != nil {
		log.Println("Failed to delete account. Error", err)
		return &types.NewError{Message: "Failed to delete account"}
	}
	return nil
}
