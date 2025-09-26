package api

import (
	"net/http"

	"github.com/Kranold/hyrox/internal/strava"
)

func (cfg *APIConfig) CreateStravaWebhookSubscription(w http.ResponseWriter, r *http.Request) {

	err := strava.CreateSubscription()

	if err != nil {
		http.Error(w, "Error creating strava webhook subscription", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

}
