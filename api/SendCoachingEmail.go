package api

import (
	"fmt"
	"net/http"

	"github.com/Kranold/hyrox/internal/aiservice"
	"github.com/Kranold/hyrox/internal/email"

	"github.com/google/uuid"
)

func (cfg *APIConfig) SendCoachingEmail(w http.ResponseWriter, r *http.Request) {
	userID, _ := r.Context().Value("userID").(uuid.UUID)

	aiservice := aiservice.AIService{
		DB: cfg.DB,
	}
	coachingAdvice, err := aiservice.AICoaching(r.Context(), userID)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error getting coaching advice", http.StatusInternalServerError)
		return
	}

	user, err := cfg.DB.GetUserByID(r.Context(), userID)
	if err != nil {
		http.Error(w, "Error getting user", http.StatusInternalServerError)
		return
	}

	err = email.SendCoachingEmail(user.Username, user.Email, coachingAdvice)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error sending email", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
