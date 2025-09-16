-- +goose up
CREATE TABLE strava_activities (
    id BIGINT PRIMARY KEY, -- Unique identifier for the activity
    external_id VARCHAR(255), -- External ID of the activity
    upload_id BIGINT, -- Upload ID of the activity
    athlete_id BIGINT REFERENCES strava_user(strava_id) ON DELETE CASCADE, -- ID of the athlete
    name VARCHAR(255), -- Name of the activity
    description TEXT, -- Description of the activity
    distance FLOAT, -- Distance covered in meters
    moving_time INT, -- Moving time in seconds
    elapsed_time INT, -- Total elapsed time in seconds
    total_elevation_gain FLOAT, -- Total elevation gain in meters
    elev_high FLOAT, -- Highest elevation in meters
    elev_low FLOAT, -- Lowest elevation in meters
    type VARCHAR(50), -- Type of activity (e.g., Run, Ride, etc.)
    sport_type VARCHAR(50), -- Sport type (e.g., MountainBikeRide, GravelRide, etc.)
    start_date VARCHAR(255), -- Start date and time in UTC
    start_date_local VARCHAR(255), -- Start date and time in local timezone
    timezone VARCHAR(100), -- Timezone of the activity
    average_speed FLOAT, -- Average speed in meters per second
    max_speed FLOAT, -- Maximum speed in meters per second
    average_cadence FLOAT, -- Average cadence (if applicable)
    average_heartrate FLOAT, -- Average heart rate (if applicable)
    max_heartrate FLOAT, -- Maximum heart rate (if applicable)
    calories FLOAT, -- Calories burned (if available)
    workout_type INT, -- Type of workout (if applicable)
    kudos_count INT, -- Number of kudos received
    comment_count INT, -- Number of comments
    achievement_count INT, -- Number of achievements
    photo_count INT, -- Number of photos
    trainer BOOLEAN, -- Whether the activity was performed on a trainer
    commute BOOLEAN, -- Whether the activity was a commute
    manual BOOLEAN, -- Whether the activity was manually entered
    private BOOLEAN, -- Whether the activity is private
    visibility VARCHAR(50), -- Visibility setting of the activity
    flagged BOOLEAN, -- Whether the activity is flagged
    gear_id VARCHAR(50), -- ID of the gear used
    splits text, -- JSONB, -- JSON array of splits (if available)
    map_summary TEXT, -- Summary polyline of the route
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Record creation timestamp
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Record update timestamp
);

-- +goose down
DROP TABLE strava_activities;
