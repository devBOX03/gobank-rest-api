package main

import (
	"log"
	"log/slog"
	"net/http"
)

type APIServer struct {
	listerAddress string
	store         Storage
}

func NewAPIServer(listerAddress string, store Storage) APIServer {
	return APIServer{
		listerAddress: listerAddress,
		store:         store,
	}
}

func (s APIServer) Run() {
	http.HandleFunc("GET /hello", s.helloHandler)
	http.HandleFunc("POST /account", s.createAccount)
	http.HandleFunc("GET /account", s.getAccounts)
	http.HandleFunc("GET /account/{id}", s.getAccountById)
	http.HandleFunc("DELETE /account/{id}", s.deleteAccountById)

	slog.Info("Server is running on port 3000 ...")
	if err := http.ListenAndServe(s.listerAddress, nil); err != nil {
		log.Fatal("Encountered error while running server: ", err)
	}
}
