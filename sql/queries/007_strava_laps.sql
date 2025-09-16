-- name: CreateStravaLap :exec
INSERT INTO strava_laps (
    id,
    activity_id,
    athlete_id,
    average_cadence,
    average_speed,
    average_heartrate,
    max_heartrate,
    distance,
    elapsed_time,
    start_index,
    end_index,
    lap_index,
    max_speed,
    moving_time,
    name,
    pace_zone,
    split,
    start_date,
    start_date_local,
    total_elevation_gain
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20
);
