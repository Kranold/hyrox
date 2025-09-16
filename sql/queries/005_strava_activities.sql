-- name: CreateStravaActivity :exec
INSERT INTO strava_activities (
    id,
    external_id,
    upload_id,
    athlete_id,
    name,
    description,
    distance,
    moving_time,
    elapsed_time,
    total_elevation_gain,
    elev_high,
    elev_low,
    type,
    sport_type,
    start_date,
    start_date_local,
    timezone,
    average_speed,
    max_speed,
    average_cadence,
    average_heartrate,
    max_heartrate,
    calories,
    workout_type,
    kudos_count,
    comment_count,
    achievement_count,
    photo_count,
    trainer,
    commute,
    manual,
    private,
    visibility,
    flagged,
    gear_id,
    splits,
    map_summary
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37
);

-- name: GetStravaActivitiesByAthleteID :many
SELECT * FROM strava_activities
WHERE athlete_id = $1
ORDER BY start_date DESC;

-- name: GetStravaActivitiesByUserID :many
SELECT sa.* FROM strava_activities sa
JOIN strava_user su ON su.strava_id = sa.athlete_id
WHERE su.user_id = $1
ORDER BY sa.start_date DESC;

-- name: GetAllStravaActivitySegmentsAndLapsByUserID :many
SELECT sa.*, ss.*, sl.* FROM strava_activities sa
LEFT JOIN strava_segments ss on ss.activity_id = sa.id 
LEFT JOIN strava_laps sl on sl.activity_id = sa.id
JOIN strava_user su ON su.strava_id = sa.athlete_id
WHERE su.user_id  = $1;
