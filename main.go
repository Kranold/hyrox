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
		DB:        *dbQueries,
		JWTSecret: os.Getenv("JWT_SECRET"),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/create_user", apiCfg.CreateUser)
	mux.Handle("/getactivities", apiCfg.AuthMiddleware(http.HandlerFunc(apiCfg.GetAllStravaActivites)))
	mux.HandleFunc("/strava_token_exchange", apiCfg.StravaTokenExchangeHandler)
	mux.HandleFunc("/delete_users", apiCfg.DeleteAllUsers)
	mux.HandleFunc("/login", apiCfg.Login)

	muxHandlerWithCORS := api.CORSMiddleware(mux)
	port := "8080"
	server := &http.Server{
		Addr:    ":" + port,
		Handler: muxHandlerWithCORS,
	}
	log.Printf("Starting on port: %s\n", port)
	log.Fatal(server.ListenAndServe())
}
