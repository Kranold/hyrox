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

		err := cfg.DB.CreateStravaActivity(ctx, database.CreateStravaActivityParams{
			ID:                 a.ID,
			ExternalID:         helperfunctions.ToNullString(a.ExternalID),
			UploadID:           helperfunctions.ToNullInt64(a.UploadID),
			AthleteID:          helperfunctions.ToNullInt64(a.Athlete.ID),
			Name:               helperfunctions.ToNullString(a.Name),
			Description:        helperfunctions.ToNullString(a.Description),
			Distance:           helperfunctions.ToNullFloat64(a.Distance),
			MovingTime:         helperfunctions.ToNullInt32FromInt(a.MovingTime),
			ElapsedTime:        helperfunctions.ToNullInt32FromInt(a.ElapsedTime),
			TotalElevationGain: helperfunctions.ToNullFloat64(a.TotalElevationGain),
			ElevHigh:           helperfunctions.ToNullFloat64(a.ElevHigh),
			ElevLow:            helperfunctions.ToNullFloat64(a.ElevLow),
			Type:               helperfunctions.ToNullString(a.Type),
			SportType:          helperfunctions.ToNullString(a.SportType),
			StartDate:          helperfunctions.ToNullString(a.StartDate),
			AverageSpeed:       helperfunctions.ToNullFloat64(a.AverageSpeed),
			MaxSpeed:           helperfunctions.ToNullFloat64(a.MaxSpeed),
			AverageCadence:     helperfunctions.ToNullFloat64(a.AverageCadence),
			AverageHeartrate:   helperfunctions.ToNullFloat64(a.AverageHeartrate),
			MaxHeartrate:       helperfunctions.ToNullFloat64(a.MaxHeartrate),
			//	Splits:             helperfunctions.ToNullString(a.Splits), // a.SplitsMetric,
			Calories: helperfunctions.ToNullFloat64(a.Calories),
		})
		if err != nil {
			fmt.Println(err)
		}
		b, err := GetDetailedStravaActivity(ctx, accessTokens.AccessToken, a.ID)
		if err != nil {
			fmt.Println(err)
		}
		if len(b.Laps) > 0 {
			for _, lap := range b.Laps {

				err := cfg.DB.CreateStravaLap(ctx, database.CreateStravaLapParams{
					ID:               lap.ID,
					ActivityID:       helperfunctions.ToNullInt64(a.ID),
					AthleteID:        helperfunctions.ToNullInt64(a.Athlete.ID),
					AverageCadence:   helperfunctions.ToNullFloat64(lap.AverageCadence),
					AverageSpeed:     helperfunctions.ToNullFloat64(lap.AverageSpeed),
					AverageHeartrate: helperfunctions.ToNullFloat64(lap.AverageHeartrate),
					MaxHeartrate:     helperfunctions.ToNullFloat64(lap.MaxHeartrate),
					Distance:         helperfunctions.ToNullFloat64(lap.Distance),
					ElapsedTime:      helperfunctions.ToNullInt32(lap.ElapsedTime),
					StartIndex:       helperfunctions.ToNullInt32(lap.StartIndex),
					EndIndex:         helperfunctions.ToNullInt32(lap.EndIndex),
					LapIndex:         helperfunctions.ToNullInt32(lap.LapIndex),
					MaxSpeed:         helperfunctions.ToNullFloat64(lap.MaxSpeed),
					MovingTime:       helperfunctions.ToNullInt32(lap.MovingTime),
					Name:             helperfunctions.ToNullString(lap.Name),
					PaceZone:         helperfunctions.ToNullInt32(lap.PaceZone),
					Split:            helperfunctions.ToNullInt32(lap.Split),
					//		StartDate:          lap.StartDate,
					//		StartDateLocal:     lap.StartDateLocal,
					TotalElevationGain: helperfunctions.ToNullFloat64(lap.TotalElevationGain),
				})
				if err != nil {
					fmt.Println(err)
				}
			}
		}
		if len(b.Segments) > 0 {
			for _, segment := range b.Segments {
				err := cfg.DB.CreateStravaSegment(ctx, database.CreateStravaSegmentParams{
					ID:          segment.ID,
					ActivityID:  helperfunctions.ToNullInt64(a.ID),
					ElapsedTime: helperfunctions.ToNullInt32(segment.ElapsedTime),
					//StartDate:        helperfunctions.ToNullString(segment.StartDate),
					// StartDateLocal:   helperfunctions.ToNullString(segment.StartDateLocal),
					Distance:         helperfunctions.ToNullFloat64(segment.Distance),
					MovingTime:       helperfunctions.ToNullInt32(segment.MovingTime),
					StartIndex:       helperfunctions.ToNullInt32(segment.StartIndex),
					EndIndex:         helperfunctions.ToNullInt32(segment.EndIndex),
					AverageCadence:   helperfunctions.ToNullFloat64(segment.AverageCadence),
					AverageWatts:     helperfunctions.ToNullFloat64(segment.AverageWatts),
					AverageHeartrate: helperfunctions.ToNullFloat64(segment.AverageHeartrate),
					MaxHeartrate:     helperfunctions.ToNullFloat64(segment.MaxHeartrate),
				})
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
	return nil
}
