package api

import (
	"encoding/json"
	"net/http"
)

func (cfg *APIConfig) ValidateStravaWebhookSubscription(w http.ResponseWriter, r *http.Request) {

	hubChallenge := r.URL.Query().Get("hub.challenge")
	hubVerifyToken := r.URL.Query().Get("hub.verify_token")

	if hubVerifyToken != "HYROX_APPLICATION_TOKEN" {
		http.Error(w, "Invalid verify token", http.StatusBadRequest)
		return
	}

	type response struct {
		HubChallenge string `json:"hub.challenge"`
	}
	respData := response{
		HubChallenge: hubChallenge,
	}

	jsonResp, _ := json.Marshal(respData)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)

}
