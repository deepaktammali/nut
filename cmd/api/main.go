package main

import (
	"fmt"
	"log"
	"net/http"
	"nut/internal/config"
	"nut/internal/database"
	"nut/internal/handlers"
)

const apiPort = "8080"

func main() {
	db, err := database.NewDb(database.DatabaseDriverPgx, database.BuildPostgresDSN())

	if err != nil {
		log.Fatalf("Error connecting to the database - %s", err)
	}

	app := config.AppConfig{
		Db:          db,
		Environment: "production",
	}

	server := http.Server{
		Addr:    fmt.Sprintf("%s", apiPort),
		Handler: handlers.NewHandler(&app),
	}

	log.Fatal(server.ListenAndServe())
}
