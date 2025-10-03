package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Kranold/hyrox/internal/database"
	"github.com/Kranold/hyrox/internal/logging"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func TestCreateUser(t *testing.T) {

	tests := map[string]struct {
		payload      map[string]interface{}
		expectedCode int
	}{
		"ValidInput": {
			payload: map[string]interface{}{
				"Email":       "testuser@gmail.com",
				"UserName":    "testuser",
				"Password":    "securepassword",
				"FitnessGoal": "Build Muscle",
				"Birthday":    "1990-01-01",
			},
			expectedCode: 201,
		},
		"AnotherValidInput": {
			payload: map[string]interface{}{
				"Email":       "test1234",
				"UserName":    "testuser123",
				"Password":    "",
				"FitnessGoal": "Build Muscle",
				"Birthday":    "1990-01-01",
			},
			expectedCode: 201,
		},
	}

	logger := logging.CreateLogger()

	/*Here is a hack to get the test running both via github actions
	and locally by loading env variables. Not pretty
	*/
	err := godotenv.Load()
	if err != nil {
		godotenv.Load("../.env")
	}

	dbURL := os.Getenv("DB_URL")

	fmt.Println("dburl:", dbURL)
	db, err := sql.Open("postgres", dbURL)

	if err != nil {
		logger.Error("Error connecting to the database", slog.String("Error", err.Error()))
		log.Fatalf("Error connecting to the database, shutting down application")
	}

	dbQueries := database.New(db)
	cfg := &APIConfig{DB: *dbQueries}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			reqBody, _ := json.Marshal(tt.payload)

			req := httptest.NewRequest("POST", "/create_user", bytes.NewReader(reqBody))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			handler := http.HandlerFunc(cfg.CreateUser)
			handler.ServeHTTP(w, req)

			res := w.Result()
			defer res.Body.Close()

			if w.Code != tt.expectedCode {
				t.Errorf("expected status %d, got %d", tt.expectedCode, w.Code)
			}

			// Cleaning the database and verifying the user was deleted

			var user database.User
			_ = json.NewDecoder(res.Body).Decode(&user)
			_ = cfg.DB.DeleteUserByID(req.Context(), user.ID)

			user, err = cfg.DB.GetUserbyUserID(req.Context(), user.ID)
			if err == nil {
				t.Errorf("Could not clean up the database, and user with ID %v still exists", user.ID)
			}

		})

	}
}
