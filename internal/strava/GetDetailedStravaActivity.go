package strava

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetDetailedStravaActivity(ctx context.Context, accessToken string, activityID int64) (StravaActivity, error) {
	activityURL := fmt.Sprintf("https://www.strava.com/api/v3/activities/%d/?include_all_efforts=", activityID)

	// Preparing the HTTP Request
	req, err := http.NewRequest("GET", activityURL, nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return StravaActivity{}, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	// Making the HTTP Request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return StravaActivity{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Unexpected status code: %d\n", resp.StatusCode)
		return StravaActivity{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Parsing the response
	var detailedActivity StravaActivity
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&detailedActivity)
	if err != nil {
		fmt.Printf("Error decoding response: %v\n", err)
		return StravaActivity{}, err
	}
	return detailedActivity, nil

}
