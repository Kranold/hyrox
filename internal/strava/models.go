package strava

import (
	"github.com/Kranold/hyrox/internal/database"
)

type StravaService struct {
	DB database.Queries
}

type StravaAthlete struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
	Premium  bool   `json:"premium"`
}

type StravaTokenResponse struct {
	TokenType    string        `json:"token_type"`
	ExpiresAt    int           `json:"expires_at"`
	ExpiresIn    int           `json:"expires_in"`
	RefreshToken string        `json:"refresh_token"`
	AccessToken  string        `json:"access_token"`
	Athlete      StravaAthlete `json:"athlete"`
}

type StravaActivity struct {
	ID                 int64     `json:"id"`
	ExternalID         string    `json:"external_id"`
	UploadID           int64     `json:"upload_id"`
	AthleteID          int64     `json:"athlete_id"`
	Name               string    `json:"name"`
	Description        string    `json:"description"`
	Distance           float64   `json:"distance"`
	MovingTime         int       `json:"moving_time"`
	ElapsedTime        int       `json:"elapsed_time"`
	TotalElevationGain float64   `json:"total_elevation_gain"`
	ElevHigh           float64   `json:"elev_high"`
	ElevLow            float64   `json:"elev_low"`
	Type               string    `json:"type"`
	SportType          string    `json:"sport_type"`
	StartDate          string    `json:"start_date"`       // ISO 8601 format (e.g., "2023-10-01T06:30:00Z")
	StartDateLocal     string    `json:"start_date_local"` // ISO 8601 format
	Timezone           string    `json:"timezone"`
	StartLatlng        []float64 `json:"start_latlng"` // Format: "lat,lng"
	EndLatlng          []float64 `json:"end_latlng"`   // Format: "lat,lng"
	AverageSpeed       float64   `json:"average_speed"`
	MaxSpeed           float64   `json:"max_speed"`
	AverageCadence     float64   `json:"average_cadence"`
	AverageHeartrate   float64   `json:"average_heartrate"`
	MaxHeartrate       float64   `json:"max_heartrate"`
	Calories           float64   `json:"calories"`
	WorkoutType        int       `json:"workout_type"`
	KudosCount         int       `json:"kudos_count"`
	CommentCount       int       `json:"comment_count"`
	AchievementCount   int       `json:"achievement_count"`
	PhotoCount         int       `json:"photo_count"`
	Trainer            bool      `json:"trainer"`
	Commute            bool      `json:"commute"`
	Manual             bool      `json:"manual"`
	Private            bool      `json:"private"`
	Visibility         string    `json:"visibility"`
	Flagged            bool      `json:"flagged"`
	GearID             string    `json:"gear_id"`
	MapSummary         string    `json:"map_summary"`
	CreatedAt          string    `json:"created_at"` // ISO 8601 format
	UpdatedAt          string    `json:"updated_at"` // ISO 8601 format
}
