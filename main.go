package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/Kranold/hyrox/api"
	"github.com/Kranold/hyrox/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()

	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	dbQueries := database.New(db)

	apiCfg := &api.APIConfig{
		DB: *dbQueries,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/create_user", apiCfg.CreateUser)
	mux.HandleFunc("/create_user", apiCfg.LinkStravaAccountToUser)

	port := "8080"
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
	log.Printf("Starting on port: %s\n", port)
	log.Fatal(server.ListenAndServe())
}
