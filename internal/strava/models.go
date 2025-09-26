package strava

import (
	"github.com/Kranold/hyrox/internal/database"
)

type StravaService struct {
	DB database.Queries
}

type StravaAthlete struct {
	ID        int     `json:"id"`
	UserName  string  `json:"username"`
	Premium   bool    `json:"premium"`
	FirstName string  `json:"firstname"`
	LastName  string  `json:"lastname"`
	City      string  `json:"city"`
	State     string  `json:"state"`
	Country   string  `json:"country"`
	Sex       string  `json:"sex"`
	Weight    float64 `json:"weight"`
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
	Athlete            Athlete   `json:"athlete"`
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
	//	Splits             string          `json:"splits_metric"`
	Segments []StravaSegment `json:"segment_efforts"`
	Laps     []StravaLap     `json:"laps"`
}

type Athlete struct {
	ID             int64 `json:"id"`
	Resource_state int   `json:"resource_state"`
}

type Activity struct {
	ID             int64 `json:"id"`
	Resource_state int   `json:"resource_state"`
}

type StravaSegment struct {
	ID               int64   `json:"id"`
	ActivityID       int64   `json:"activity_id"`
	ElapsedTime      int32   `json:"elapsed_time"`
	StartDate        string  `json:"start_date"`
	StartDateLocal   string  `json:"start_date_local"`
	Distance         float64 `json:"distance"`
	MovingTime       int32   `json:"moving_time"`
	StartIndex       int32   `json:"start_index"`
	EndIndex         int32   `json:"end_index"`
	AverageCadence   float64 `json:"average_cadence"`
	AverageWatts     float64 `json:"average_watts"`
	AverageHeartrate float64 `json:"average_heartrate"`
	MaxHeartrate     float64 `json:"max_heartrate"`
}

type StravaLap struct {
	ID                 int64    `json:"id"`
	Activity           Activity `json:"activity_id"`
	Athlete            Athlete  `json:"athlete_id"`
	AverageCadence     float64  `json:"average_cadence"`
	AverageSpeed       float64  `json:"average_speed"`
	AverageHeartrate   float64  `json:"average_heartrate"`
	MaxHeartrate       float64  `json:"max_heartrate"`
	Distance           float64  `json:"distance"`
	ElapsedTime        int32    `json:"elapsed_time"`
	StartIndex         int32    `json:"start_index"`
	EndIndex           int32    `json:"end_index"`
	LapIndex           int32    `json:"lap_index"`
	MaxSpeed           float64  `json:"max_speed"`
	MovingTime         int32    `json:"moving_time"`
	Name               string   `json:"name"`
	PaceZone           int32    `json:"pace_zone"`
	Split              int32    `json:"split"`
	StartDate          string   `json:"start_date"`
	StartDateLocal     string   `json:"start_date_local"`
	TotalElevationGain float64  `json:"total_elevation_gain"`
}

type StravaEventHookParameters struct {
	ObjectType string `json:"object_type"`
	ObjectID   int64  `json:"object_id"`
	AspectType string `json:"aspect_type"`
	//	Updates        string `json:"updates"`
	OwnerID        int64 `json:"owner_id"`
	SubscriptionID int64 `json:"subscription_id"`
	EventTime      int64 `json:"event_time"`
}
