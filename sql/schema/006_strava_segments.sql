-- +goose up 
CREATE TABLE strava_segments(
    id BIGINT PRIMARY KEY, -- Unique identifier for the segment
    activity_id BIGINT REFERENCES strava_activities(id) ON DELETE CASCADE, -- ID of the associated activity
    elapsed_time INT, -- Total elapsed time in seconds
    start_date VARCHAR(255), -- Start date and time of the segment in UTC
    start_date_local VARCHAR(255), -- Start date and time of the segment in local timezone
    distance FLOAT, -- Distance covered in meters
    moving_time INT, -- Moving time in seconds
    start_index INT, -- Start index of the segment within the activity
    end_index INT, -- End index of the segment within the activity
    average_cadence FLOAT, -- Average cadence (if applicable)
    average_watts FLOAT, -- Average power output in watts (if applicable)
    average_heartrate FLOAT, -- Average heart rate (if applicable)
    max_heartrate FLOAT -- Maximum heart rate (if applicable)
 );

-- +goose down
DROP TABLE strava_segments;   