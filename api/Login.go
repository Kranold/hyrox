package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Kranold/hyrox/internal/auth"
	"github.com/Kranold/hyrox/internal/database"
)

func (cfg *APIConfig) Login(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	loginParams := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&loginParams)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user, err := cfg.DB.GetUserByEmail(r.Context(), loginParams.Email)
	if err != nil {
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
			fmt.Println(err)
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
			fmt.Println(err)
			http.Error(w, "Error updating refresh token", http.StatusInternalServerError)
			return

		}
	}

	//preparing the response
	type respData struct {
		JWT          string                `json:"jwt"`
		RefreshToken database.RefreshToken `json:"refresh_token"`
		User         database.User         `json:"user"`
	}

	userRespData := database.User{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	responseData := respData{
		JWT:          jwtToken,
		RefreshToken: refreshTokenData,
		User:         userRespData,
	}

	data, err := json.Marshal(responseData)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}

	// Sending the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
