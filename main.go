package main

import (
	"log"
)

func main() {
	db, err := newDB()
	if err != nil {
		log.Fatal("Failed to connect to DB due to error: ", err)
		panic("Failed to connect to DB")
	}
	defer db.Close()
	store := NewStorage(db)

	server := NewAPIServer(":3000", store)
	server.Run()

}
