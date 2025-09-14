package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Kranold/hyrox/internal/auth"
	"github.com/Kranold/hyrox/internal/database"
	"github.com/Kranold/hyrox/internal/helperfunctions"
	"github.com/Kranold/hyrox/internal/strava"
)

func (cfg *APIConfig) StravaTokenExchangeHandler(w http.ResponseWriter, r *http.Request) {
	authcode := r.URL.Query().Get("code")
	jwtToken := r.URL.Query().Get("token")

	if authcode == "" {
		http.Error(w, "Missing auth code", http.StatusBadRequest)
		return
	}
	// Exchange auth code for access and refresh tokens
	tokenResponse, err := strava.GetStravaAccessTokens(authcode)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error exchanging auth code for tokens", http.StatusInternalServerError)
		return
	}

	//Store tokens and stravaData in db

	userID, err := auth.ValidateJWT(jwtToken, cfg.JWTSecret)
	if err != nil {
		http.Error(w, "Invalid JWT token", http.StatusUnauthorized)
		return
	}
	newStravaUserData := database.CreateStravaUserParams{
		UserID:    userID,
		StravaID:  int64(tokenResponse.Athlete.ID),
		Username:  tokenResponse.Athlete.UserName,
		Firstname: helperfunctions.ToNullString(tokenResponse.Athlete.FirstName),
		Lastname:  helperfunctions.ToNullString(tokenResponse.Athlete.LastName),
		City:      helperfunctions.ToNullString(tokenResponse.Athlete.City),
		State:     helperfunctions.ToNullString(tokenResponse.Athlete.State),
		Country:   helperfunctions.ToNullString(tokenResponse.Athlete.Country),
		Sex:       helperfunctions.ToNullString(tokenResponse.Athlete.Sex),
		Premuim:   helperfunctions.ToNullBool(tokenResponse.Athlete.Premium),
		Weight:    helperfunctions.ToNullInt32(int(tokenResponse.Athlete.Weight)),
	}

	stravaUser, err := cfg.DB.CreateStravaUser(r.Context(), newStravaUserData)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error storing strava user data", http.StatusInternalServerError)
		return
	}

	newStravaAccessToken := database.CreateStravaAccessTokenParams{
		StravaUserID: stravaUser.ID,
		AccessToken:  tokenResponse.AccessToken,
		TokenType:    tokenResponse.TokenType,
		ExpiresAt:    time.Unix(int64(tokenResponse.ExpiresAt), 0),
		RefreshToken: tokenResponse.RefreshToken,
	}

	_, err = cfg.DB.CreateStravaAccessToken(r.Context(), newStravaAccessToken)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error storing strava access token", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
