package api

import (
	"database/sql"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/Kranold/hyrox/internal/auth"
	"github.com/Kranold/hyrox/internal/database"
	"github.com/Kranold/hyrox/internal/helperfunctions"
	"github.com/Kranold/hyrox/internal/logging"
)

func (cfg *APIConfig) CreateUser(w http.ResponseWriter, r *http.Request) {
	logger := logging.CreateLogger()
	// Parse the JSON request body
	type parameters struct {
		Email       string `json:"email"`
		UserName    string `json:"username"`
		Password    string `json:"password"`
		FitnessGoal string `json:"fitness_goal"`
		Birthday    string `json:"birthday"` //expecting YYY-MM-DD
	}
	newUser := parameters{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		logger.Error("Error decoding request body",
			slog.String("Error", err.Error()))
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	hashedPassword, _ := auth.HashPassword(newUser.Password)
	// Creating the user
	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		Email:          newUser.Email,
		Username:       newUser.UserName,
		FitnessGoal:    helperfunctions.ToNullString(newUser.FitnessGoal),
		Birthday:       helperfunctions.StringDateToDate(newUser.Birthday),
		HashedPassword: sql.NullString{String: hashedPassword, Valid: true},
	})
	if err != nil {
		logger.Error("Error creating user",
			slog.Any("User", newUser),
			slog.String("Error", err.Error()))
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	// Preparing the response
	respData := database.User{
		ID:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		FitnessGoal: user.FitnessGoal,
		Birthday:    user.Birthday,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}

	data, err := json.Marshal(respData)
	if err != nil {
		logger.Error("Error encoding response",
			slog.String("Error", err.Error()))
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}

	// Sending the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(data)

}
