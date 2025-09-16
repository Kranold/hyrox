-- name: CreateStravaSegment :exec
INSERT INTO strava_segments (
    id,
    activity_id,
    elapsed_time,
    start_date,
    start_date_local,
    distance,
    moving_time,
    start_index,
    end_index,
    average_cadence,
    average_watts,
    average_heartrate,
    max_heartrate   
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
);