package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, logger *slog.Logger, code int, payload interface{}) {

	data, err := json.Marshal(payload)
	if err != nil {
		logger.Error("Error encoding JSON response",
			slog.String("Error", err.Error()))
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}

	// Sending the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}
