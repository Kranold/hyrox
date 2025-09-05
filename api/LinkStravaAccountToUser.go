package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/Kranold/hyrox/internal/database"
	"github.com/Kranold/hyrox/internal/strava"
	"github.com/google/uuid"
)

func (cfg *APIConfig) LinkStravaAccountToUser(w http.ResponseWriter, r *http.Request) {

	// Parsing the the request
	type parameters struct {
		UserID   string `json:"user_id"`
		AuthCode string `json:"auth_code"`
	}
	reqParams := parameters{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqParams)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	//Internal logic

	// Getting Strava access tokens and athlete data
	stravaTokens, err := strava.GetStravaAccessTokens(reqParams.AuthCode)
	if err != nil {
		http.Error(w, "Error getting Strava access tokens", http.StatusInternalServerError)
		return
	}

	// Creating new strava row
	athlete, err := strava.GetStravaAthlete(stravaTokens.AccessToken)
	if err != nil {
		http.Error(w, "Error getting Strava athlete data", http.StatusInternalServerError)
		return
	}

	userID, _ := uuid.Parse(reqParams.UserID)
	stravaUserData := database.CreateStravaUserParams{
		UserID:                userID,
		StravaID:              int64(athlete.ID),
		RefreshToken:          stravaTokens.RefreshToken,
		RefreshTokenExpiresAt: time.Unix(int64(stravaTokens.ExpiresAt), 0),
		Username:              athlete.UserName,
		Premuim:               sql.NullBool{Bool: athlete.Premium, Valid: true},
	}

	stravaUser, err := cfg.DB.CreateStravaUser(r.Context(), stravaUserData)
	if err != nil {
		http.Error(w, "Error creating Strava athlete data", http.StatusInternalServerError)
		return
	}

	//preparing the response
	respData := database.StravaUser{
		UserID:    stravaUser.UserID,
		StravaID:  stravaUser.StravaID,
		Username:  stravaUser.Username,
		CreatedAt: stravaUser.CreatedAt,
		UpdatedAt: stravaUser.UpdatedAt,
		Premuim:   stravaUser.Premuim,
	}
	data, err := json.Marshal(respData)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}

	//Sending the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(data)

}
