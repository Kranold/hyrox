package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/Kranold/hyrox/internal/aiservice"
	"github.com/Kranold/hyrox/internal/database"
	email "github.com/Kranold/hyrox/internal/emailservice"
	"github.com/Kranold/hyrox/internal/logging"
	"github.com/Kranold/hyrox/internal/strava"
)

// https://developers.strava.com/docs/webhooks/

type stravaEventHookParameters struct {
	ObjectType string `json:"object_type"`
	ObjectID   int64  `json:"object_id"`
	AspectType string `json:"aspect_type"`
	//	Updates        string `json:"updates"`
	OwnerID        int64 `json:"owner_id"`
	SubscriptionID int64 `json:"subscription_id"`
	EventTime      int64 `json:"event_time"`
}

func (cfg *APIConfig) StravaWebhookHandler(w http.ResponseWriter, r *http.Request) {
	logger := logging.CreateLogger()

	logger.Info("Received Strava webhook event")

	eventParams := stravaEventHookParameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&eventParams)
	if err != nil {
		logger.Error("Error decoding strava webhook event",
			slog.String("Error", err.Error()))
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	logger.Debug("Strava webhook event parameters",
		slog.String("ObjectType", eventParams.ObjectType))
	// Handle activity create

	if eventParams.ObjectType == "activity" && eventParams.AspectType == "create" {
		_, err := cfg.DB.GetStravaActivityByID(r.Context(), eventParams.ObjectID)
		if err == nil {
			// Activity already exists, no need to process further
			return
		}

		go func() {
			err = HandleActivityCreate(context.Background(), cfg.DB, eventParams)
			if err != nil {
				logger.Error("Error handling strava activity create",
					slog.String("Error", err.Error()))
				return
			}
		}()
		return
	}

	// TODO handle activity update
	// TODO handle activity delete
	// TODO handle athlete update
	// TODO handle athlete delete
	w.WriteHeader(200)
	logger.Info("Processed strava webhook event",
		slog.String("ObjectType", eventParams.ObjectType),
		slog.String("AspectType", eventParams.AspectType),
		slog.Int64("ObjectID", eventParams.ObjectID),
		slog.Int64("OwnerID", eventParams.OwnerID))

}

func HandleActivityCreate(ctx context.Context, db database.Queries, eventParams stravaEventHookParameters) error {
	logger := logging.CreateLogger()
	StravaService := strava.StravaService{
		DB: db,
	}
	aiservice := aiservice.AIService{
		DB: db,
	}
	client := &http.Client{}

	user, err := db.GetUserFromStravaID(ctx, eventParams.OwnerID)
	if err != nil {
		logger.Error("Error getting user for stravaID",
			slog.Int64("StravaID", eventParams.OwnerID),
			slog.String("Error", err.Error()))
		return err
	}

	accessToken, _ := StravaService.GetStravaAccessToken(ctx, client, user.ID)
	if accessToken.AccessToken == "" {
		logger.Error("No access token found for user",
			slog.Int64("StravaID", eventParams.OwnerID),
			slog.Any("UserID", user.ID))
		return fmt.Errorf("no access token found for user with StravaID %d", eventParams.OwnerID)
	}

	activity, err := strava.GetStravaActivity(ctx, client, accessToken.AccessToken, eventParams.ObjectID)
	if err != nil {
		logger.Error("Error getting detailed strava activity",
			slog.Int64("ActivityID", eventParams.ObjectID),
			slog.String("Error", err.Error()))
		return err
	}

	err = StravaService.SaveStravaActivity(ctx, activity)
	if err != nil {
		logger.Error("Error saving detailed strava activity",
			slog.Int64("ActivityID", eventParams.ObjectID),
			slog.String("Error", err.Error()))
		return err
	}

	// Getting AI coaching advice and sending email
	coachingAdvice, err := aiservice.AICoaching(ctx, user.ID)
	if err != nil {
		logger.Error("Error getting AI coaching advice",
			slog.String("Error", err.Error()),
			slog.String("UserID", user.ID.String()))
		return err
	}
	err = email.SendCoachingEmail(user.Username, user.Email, coachingAdvice)
	if err != nil {
		logger.Error("Error sending email with AI coaching advice",
			slog.String("Error", err.Error()),
			slog.String("Email", user.Email),
			slog.String("UserID", user.ID.String()))
		return err
	}

	return nil
}
