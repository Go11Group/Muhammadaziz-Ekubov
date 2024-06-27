package main

import (
	postgres "atto/storage/postgres"
	"log"
)

func main() {
	db, err := postgres.CreateDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

}
