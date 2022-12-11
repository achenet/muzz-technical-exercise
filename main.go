package main

import (
	"log"
	"math/rand"
	"muzz/db"
	"muzz/handlers"
	"time"
)

func main() {
	// seed RNG to have different values each run
	rand.Seed(time.Now().UnixNano())

	database, err := db.Connect()
	if err != nil {
		log.Fatalf("failed to connect to db: %w", err)
	}

	api := handlers.NewAPI(database)

	api.Logger.Fatal(api.Start(":8080"))
}
