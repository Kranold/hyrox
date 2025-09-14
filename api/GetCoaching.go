package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Kranold/hyrox/internal/aiservice"
	"github.com/google/uuid"
)

func (cfg *APIConfig) GetCoaching(w http.ResponseWriter, r *http.Request) {
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

	type coachingAdviceJSON struct {
		Suggested_next_training string
		Focus_for_next_week     string
		Injury_preventaion      string
	}

	var jsonCoachingadvice coachingAdviceJSON

	err = json.Unmarshal([]byte(coachingAdvice), &jsonCoachingadvice)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error converting to json", http.StatusInternalServerError)
		return
	}

	data, _ := json.Marshal(jsonCoachingadvice)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
