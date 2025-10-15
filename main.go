package main

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/Kranold/hyrox/api"
	"github.com/Kranold/hyrox/internal/database"
	"github.com/Kranold/hyrox/internal/logging"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()

	logger := logging.CreateLogger()

	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		logger.Error("Error connecting to the database", slog.String("Error", err.Error()))
		log.Fatalf("Error connecting to the database, shutting down application")
	}
	dbQueries := database.New(db)

	apiCfg := &api.APIConfig{
		DB:        *dbQueries,
		JWTSecret: os.Getenv("JWT_SECRET"),
	}

	mux := http.NewServeMux()

	// serving the front-end
	mux.Handle("/", http.FileServer(http.Dir("./static")))

	// serving API calls
	mux.HandleFunc("/create_user", apiCfg.CreateUser)
	mux.Handle("/sync_strava_activities", apiCfg.AuthMiddleware(http.HandlerFunc(apiCfg.SyncStravaActivites)))
	mux.HandleFunc("/strava_token_exchange", apiCfg.StravaTokenExchangeHandler)
	mux.HandleFunc("/delete_users", apiCfg.DeleteAllUsers)
	mux.HandleFunc("/login", apiCfg.Login)
	mux.Handle("/send_coaching_email", apiCfg.AuthMiddleware(http.HandlerFunc(apiCfg.SendCoachingEmail)))
	mux.HandleFunc("/strava_webhook_create", apiCfg.CreateStravaWebhookSubscription)
	mux.HandleFunc("GET /validate_strava_subscription", apiCfg.ValidateStravaWebhookSubscription)
	mux.HandleFunc("POST /validate_strava_subscription", apiCfg.StravaWebhookHandler)

	muxHandlerWithCORS := api.CORSMiddleware(mux)
	port := os.Getenv("PORT")
	server := &http.Server{
		Addr:    ":" + port,
		Handler: muxHandlerWithCORS,
	}

	logger.Info(fmt.Sprintf("Starting on port: %s\n", port))
	log.Fatal(server.ListenAndServe())

}
