package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (app *APIServer) helloHandler(w http.ResponseWriter, r *http.Request) {
	data_map := map[string]string{
		"firstName": "Debasish",
		"lastName":  "Padhi",
		"age":       "28",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data_map)
}

func (app *APIServer) createAccount(w http.ResponseWriter, r *http.Request) {
	reqBody := new(CreateAccountRequest)
	if err := json.NewDecoder(r.Body).Decode(reqBody); err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	account, err := CreateAccountService(&app.store, reqBody.FirstName, reqBody.LastName, reqBody.Password)
	if err != nil {
		if customErr, ok := err.(*NewError); ok {
			http.Error(w, customErr.Message, http.StatusInternalServerError)
		} else {
			http.Error(w, "", http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(account)
}

func (app *APIServer) getAccounts(w http.ResponseWriter, _ *http.Request) {
	accounts, err := GetAllAccountsService(&app.store)
	if err != nil {
		if customErr, ok := err.(*NewError); ok {
			http.Error(w, customErr.Message, http.StatusInternalServerError)
		} else {
			http.Error(w, "", http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accounts)
}

func (app *APIServer) getAccountById(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}
	account, err := GetAccountByIdService(&app.store, int(id))
	if err != nil {
		if customErr, ok := err.(*NewError); ok {
			http.Error(w, customErr.Message, http.StatusInternalServerError)
		} else {
			http.Error(w, "", http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(account)
}

func (app *APIServer) deleteAccountById(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}
	sericeError := DeleteAccountByIdService(&app.store, int(id))
	if sericeError != nil {
		if customErr, ok := sericeError.(*NewError); ok {
			http.Error(w, customErr.Message, http.StatusInternalServerError)
		} else {
			http.Error(w, "", http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Successfully deleted the account")
}
