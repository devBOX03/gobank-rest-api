package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	service "github.com/devBOX03/gobank-rest-api/internal/services"
	"github.com/devBOX03/gobank-rest-api/types"
)

func (app *Application) healthHandler(w http.ResponseWriter, r *http.Request) {
	data_map := map[string]string{
		"application": "healthy",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data_map)
}

func (app *Application) createAccount(w http.ResponseWriter, r *http.Request) {
	reqBody := new(types.CreateAccountRequest)
	if err := json.NewDecoder(r.Body).Decode(reqBody); err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	account, err := service.CreateAccountService(&app.store, reqBody.FirstName, reqBody.LastName, reqBody.Password)
	if err != nil {
		if customErr, ok := err.(*types.NewError); ok {
			http.Error(w, customErr.Message, http.StatusInternalServerError)
		} else {
			http.Error(w, "", http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(account)
}

func (app *Application) getAccounts(w http.ResponseWriter, _ *http.Request) {
	accounts, err := service.GetAllAccountsService(&app.store)
	if err != nil {
		if customErr, ok := err.(*types.NewError); ok {
			http.Error(w, customErr.Message, http.StatusInternalServerError)
		} else {
			http.Error(w, "", http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accounts)
}

func (app *Application) getAccountById(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}
	account, err := service.GetAccountByIdService(&app.store, int(id))
	if err != nil {
		if customErr, ok := err.(*types.NewError); ok {
			http.Error(w, customErr.Message, http.StatusInternalServerError)
		} else {
			http.Error(w, "", http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(account)
}

func (app *Application) deleteAccountById(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}
	sericeError := service.DeleteAccountByIdService(&app.store, int(id))
	if sericeError != nil {
		if customErr, ok := sericeError.(*types.NewError); ok {
			http.Error(w, customErr.Message, http.StatusInternalServerError)
		} else {
			http.Error(w, "", http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Successfully deleted the account")
}
