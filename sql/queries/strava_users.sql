-- name: CreateStravaUser :one
INSERT INTO STRAVA_USER (id, user_id, strava_id, refresh_token, refresh_token_expires_at,
    created_at,updated_at, username, firstname, lastname, city, state, country, sex,
    premuim, weight)
VALUES (
    gen_random_uuid(),
    $1,$2,$3,$4,NOW(),NOW(),$5,$6,$7,$8,$9,$10,$11,$12,$13 
)
RETURNING  *; 

-- name: GetStravaRefreshToken :one
SELECT refresh_token,refresh_token_expires_at FROM STRAVA_USER
WHERE user_id = $1; 

-- name: UpdateStravaRefreshToken :exec
UPDATE STRAVA_USER
SET refresh_token = $2,
    refresh_token_expires_at = $3
Where user_id=$1;  