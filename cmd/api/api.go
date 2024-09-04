package main

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/devBOX03/gobank-rest-api/internal/store"
)

type Application struct {
	listerAddress string
	store         store.Storage
}

func NewAPIServer(listerAddress string, store store.Storage) Application {
	return Application{
		listerAddress: listerAddress,
		store:         store,
	}
}

func (app Application) Run() {
	http.HandleFunc("GET /health", app.healthHandler)
	http.HandleFunc("POST /account", app.createAccount)
	http.HandleFunc("GET /account", app.getAccounts)
	http.HandleFunc("GET /account/{id}", app.getAccountById)
	http.HandleFunc("DELETE /account/{id}", app.deleteAccountById)

	slog.Info("Server is running on port 3000 ...")
	if err := http.ListenAndServe(app.listerAddress, nil); err != nil {
		log.Fatal("Encountered error while running server: ", err)
	}
}
