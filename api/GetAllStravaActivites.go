package api

import (
	"log/slog"
	"net/http"

	"github.com/Kranold/hyrox/internal/logging"
	"github.com/Kranold/hyrox/internal/strava"
	"github.com/google/uuid"
)

func (cfg *APIConfig) SyncStravaActivites(w http.ResponseWriter, r *http.Request) {
	logger := logging.CreateLogger()

	userID, _ := r.Context().Value("userID").(uuid.UUID)
	// Fetching accessToken from db
	stravaService := strava.StravaService{
		DB: cfg.DB,
	}
	client := &http.Client{}
	err := stravaService.SyncStravaActivites(r.Context(), client, userID)
	if err != nil {
		logger.Error("Error fetching and saving strava activities",
			slog.String("UserID", userID.String()),
			slog.String("Error", err.Error()))
		http.Error(w, "Error fetching and saving strava activities", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}
