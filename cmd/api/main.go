package main

import (
	"log"

	"github.com/devBOX03/gobank-rest-api/internal/db"
	"github.com/devBOX03/gobank-rest-api/internal/store"
)

func main() {
	db, err := db.NewDB()
	if err != nil {
		log.Fatal("Failed to connect to DB due to error: ", err)
		panic("Failed to connect to DB")
	}
	defer db.Close()
	store := store.NewStorage(db)

	server := NewAPIServer(":3000", store)
	server.Run()
}
