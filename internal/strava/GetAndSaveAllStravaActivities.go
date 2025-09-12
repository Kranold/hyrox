package strava

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Kranold/hyrox/internal/database"
	"github.com/Kranold/hyrox/internal/helperfunctions"
	"github.com/google/uuid"
)

func (cfg *StravaService) GetAndSaveAllStravaActivites(ctx context.Context, userID uuid.UUID) error {

	// get an accces token to make the request
	accessTokens, err := cfg.GetNewStravaAccessToken(ctx, userID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	// create the request
	activityURL := "https://www.strava.com/api/v3/athlete/activities"

	req, err := http.NewRequest("GET", activityURL, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessTokens.AccessToken))

	// Send the request and extract the activities from the response

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Unexpected status code: %d\n", resp.StatusCode)
		return err
	}

	var activities []StravaActivity
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&activities)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// save all the activites to the database

	for _, a := range activities {
		fmt.Println(a)
		err := cfg.DB.CreateStravaActivity(ctx, database.CreateStravaActivityParams{
			ID:          a.ID,
			ExternalID:  helperfunctions.ToNullString(a.ExternalID),
			UploadID:    helperfunctions.ToNullInt64(a.UploadID),
			AthleteID:   helperfunctions.ToNullInt64(a.AthleteID),
			Name:        helperfunctions.ToNullString(a.Name),
			Description: helperfunctions.ToNullString(a.Description),
		})
		if err != nil {
			fmt.Println(err)
		}
	}
	return nil

}
