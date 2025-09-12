package api

import (
	"fmt"
	"net/http"

	"github.com/Kranold/hyrox/internal/strava"
	"github.com/google/uuid"
)

func (cfg *APIConfig) GetAllStravaActivites(w http.ResponseWriter, r *http.Request) {
	userID, _ := r.Context().Value("userID").(uuid.UUID)

	// Fetching accessToken from db
	stravaService := strava.StravaService{
		DB: cfg.DB,
	}

	err := stravaService.GetAndSaveAllStravaActivites(r.Context(), userID)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error fetching and saving strava activities", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

}
