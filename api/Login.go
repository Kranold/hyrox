package api

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/Kranold/hyrox/internal/auth"
	"github.com/Kranold/hyrox/internal/database"
	"github.com/Kranold/hyrox/internal/logging"
)

func (cfg *APIConfig) Login(w http.ResponseWriter, r *http.Request) {
	logger := logging.CreateLogger()

	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	loginParams := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&loginParams)
	if err != nil {
		logger.Error("Error decoding request body",
			slog.String("Error", err.Error()))
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user, err := cfg.DB.GetUserByEmail(r.Context(), loginParams.Email)
	if err != nil {
		logger.Error("Error fetching user by email",
			slog.String("Email", loginParams.Email),
			slog.String("Error", err.Error()))
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// check password and email against db
	pwCheck := auth.CheckPasswordHash(loginParams.Password, user.HashedPassword.String)

	if pwCheck != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// creating a jwt token
	expiryTime := time.Duration(360000 * time.Second)
	jwtToken, err := auth.MakeJWT(user.ID, cfg.JWTSecret, expiryTime)
	if err != nil {
		logger.Error("Error creating JWT token",
			slog.String("UserID", user.ID.String()),
			slog.String("Error", err.Error()))
		http.Error(w, "Error creating JWT token", http.StatusInternalServerError)
		return
	}

	currentTokenData, err := cfg.DB.GetRefreshToken(r.Context(), user.ID)

	var refreshTokenData database.RefreshToken

	if err != nil {
		//Create and store a local refresh token
		fmt.Printf("im here")
		refreshTokenData, err = cfg.DB.CreateRefreshToken(r.Context(), database.CreateRefreshTokenParams{
			UserID:    user.ID,
			Token:     auth.MakeRefreshToken(),
			ExpiresAt: time.Now().Add(expiryTime),
		})
		if err != nil {
			logger.Error("Error creating refresh token",
				slog.String("UserID", user.ID.String()),
				slog.String("Error", err.Error()))
			http.Error(w, "Error creating refresh token", http.StatusInternalServerError)
			return

		}

	} else {

		newTokenData := database.UpdateRefreshTokenParams{
			Token:     auth.MakeRefreshToken(),
			ExpiresAt: time.Now().Add(expiryTime),
			UserID:    user.ID,
			Token_2:   currentTokenData.Token,
		}

		refreshTokenData, err = cfg.DB.UpdateRefreshToken(r.Context(), newTokenData)

		if err != nil {
			logger.Error("Error updating refresh token",
				slog.String("UserID", user.ID.String()),
				slog.String("Error", err.Error()))
			http.Error(w, "Error updating refresh token", http.StatusInternalServerError)
			return

		}
	}

	_, err = cfg.DB.GetStravaUserByUserID(r.Context(), user.ID)

	hasStravaLinkedAccount := false

	if err == nil {
		hasStravaLinkedAccount = true
	}

	//preparing the response
	type respData struct {
		JWT                    string                `json:"jwt"`
		RefreshToken           database.RefreshToken `json:"refresh_token"`
		HasStravaLinkedAccount bool                  `json:"has_strava_linked_account"`
		User                   database.User         `json:"user"`
	}

	userRespData := database.User{
		ID:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		FitnessGoal: user.FitnessGoal,
		Birthday:    user.Birthday,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}
	responseData := respData{
		JWT:                    jwtToken,
		RefreshToken:           refreshTokenData,
		HasStravaLinkedAccount: hasStravaLinkedAccount,
		User:                   userRespData,
	}

	RespondWithJSON(w, logger, http.StatusOK, responseData)

}
