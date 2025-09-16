-- +goose up
CREATE TABLE strava_laps (
    id BIGINT PRIMARY KEY, -- Unique identifier for the lap
    activity_id BIGINT REFERENCES strava_activities(id) ON DELETE CASCADE, -- ID of the associated activity
    athlete_id BIGINT REFERENCES strava_user(strava_id) ON DELETE CASCADE, -- ID of the athlete
    average_cadence FLOAT, -- Average cadence (if applicable)
    average_speed FLOAT, -- Average speed in meters per second
    average_heartrate FLOAT, -- Average heart rate (if applicable)
    max_heartrate FLOAT, -- Maximum heart rate (if applicable)
    distance FLOAT, -- Distance covered in meters
    elapsed_time INT, -- Total elapsed time in seconds
    start_index INT, -- Start index of the lap within the activity
    end_index INT, -- End index of the lap within the activity
    lap_index INT, -- Index of the lap within the activity
    max_speed FLOAT, -- Maximum speed in meters per second
    moving_time INT, -- Moving time in seconds   
    name VARCHAR(255), -- Name of the lap
    pace_zone INT, -- Pace zone (if applicable)
    split INT, -- Split number (if applicable)
    start_date VARCHAR(255), -- Start date and time of the lap in UTC
    start_date_local VARCHAR(255), -- Start date and time of the lap in local timezone
    total_elevation_gain FLOAT, -- Total elevation gain in meters
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Record creation timestamp
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Record update timestamp
);

-- +goose down
DROP TABLE strava_laps;