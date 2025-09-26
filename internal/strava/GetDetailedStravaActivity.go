package strava

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/Kranold/hyrox/internal/logging"
)

func GetDetailedStravaActivity(ctx context.Context, accessToken string, activityID int64) (StravaActivity, error) {
	logger := logging.CreateLogger()

	activityURL := fmt.Sprintf("https://www.strava.com/api/v3/activities/%d/?include_all_efforts=", activityID)

	// Preparing the HTTP Request
	req, err := http.NewRequest("GET", activityURL, nil)
	if err != nil {
		logger.Error("Error creating request to get activity",
			slog.String("Error", err.Error()))
		return StravaActivity{}, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	// Making the HTTP Request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("Error making request to get activity",
			slog.String("Error", err.Error()))
		return StravaActivity{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logger.Error("Unexpected status code when getting activity",
			slog.Int("StatusCode", resp.StatusCode),
			slog.Any("ResponseBody", resp.Body))
		return StravaActivity{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Parsing the response
	var detailedActivity StravaActivity
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&detailedActivity)
	if err != nil {
		logger.Error("Error decoding activity response",
			slog.String("Error", err.Error()))
		return StravaActivity{}, err
	}

	return detailedActivity, nil

}
