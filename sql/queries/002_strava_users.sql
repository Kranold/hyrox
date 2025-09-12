-- name: CreateStravaUser :one
INSERT INTO strava_user (id, user_id, strava_id, created_at,updated_at, username, firstname, lastname, city, state, country, sex,
    premuim, weight)
VALUES (
    gen_random_uuid(),
    $1,$2,NOW(),NOW(),$3,$4,$5,$6,$7,$8,$9,$10,$11
)
RETURNING  *; 

