package strava

import (
	"context"
	"fmt"

	"github.com/Kranold/hyrox/internal/database"
	"github.com/Kranold/hyrox/internal/helperfunctions"
)

func (cfg *StravaService) SaveStravaActivity(ctx context.Context, a StravaActivity) error {

	// Saving the activity to the database
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
		//	Splits:             helperfunctions.ToNullString(a.Splits),
		Calories: helperfunctions.ToNullFloat64(a.Calories),
	})

	if err != nil {
		fmt.Println(err)
		return err
	}
	//Saving laps
	if len(a.Laps) > 0 {
		for _, lap := range a.Laps {

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

	//Saving segments
	if len(a.Segments) > 0 {
		for _, segment := range a.Segments {
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
	return nil
}
