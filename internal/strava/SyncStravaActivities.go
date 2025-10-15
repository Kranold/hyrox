package strava

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/Kranold/hyrox/internal/logging"
	"github.com/google/uuid"
)

func (cfg *StravaService) SyncStravaActivites(ctx context.Context, client *http.Client, userID uuid.UUID) error {
	logger := logging.CreateLogger()
	// get an accces token to make the request
	accessTokens, err := cfg.GetStravaAccessToken(ctx, client, userID)
	if err != nil {
		logger.Error("Error getting Strava access token",
			slog.String("Error", err.Error()))
		return err
	}
	// create the request
	activityURL := StravaAPIDomain + "/athlete/activities"

	req, err := http.NewRequest("GET", activityURL, nil)
	if err != nil {
		logger.Error("Error creating request to get activities",
			slog.String("Error", err.Error()))
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessTokens.AccessToken))

	// Send the request and extract the activities from the response

	resp, err := client.Do(req)
	if err != nil {
		logger.Error("Error making request to get activities",
			slog.String("Error", err.Error()))
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logger.Error("Unexpected status code when getting activities",
			slog.Int("StatusCode", resp.StatusCode),
			slog.Any("ResponseBody", resp.Body))
		return err
	}

	var activities []StravaActivity

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&activities)
	if err != nil {
		logger.Error("Error decoding activities response",
			slog.String("Error", err.Error()))
		return err
	}

	// save all the activites to the database
	for _, a := range activities {

		activity, err := GetStravaActivity(ctx, client, accessTokens.AccessToken, a.ID)
		if err != nil {
			logger.Error("Error getting strava activity",
				slog.Int64("activityID", a.ID),
				slog.String("Error", err.Error()))
			continue
		}

		err = cfg.SaveStravaActivity(ctx, activity)
		if err != nil {
			logger.Error("Error saving strava activity",
				slog.Int64("activityID", a.ID),
				slog.String("Error", err.Error()))
		}
	}

	return nil
}
