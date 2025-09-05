package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Kranold/hyrox/internal/database"
)

func (cfg *APIConfig) CreateUser(w http.ResponseWriter, r *http.Request) {

	// Parse the JSON request body
	type parameters struct {
		Email    string `json:"email"`
		UserName string `json:"username"`
	}
	newUser := parameters{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Creating the user
	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		Email:    newUser.Email,
		Username: newUser.UserName,
	})
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}

	// Preparing the ressponse
	respData := database.User{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	data, err := json.Marshal(respData)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}

	// Sending the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(data)

}
