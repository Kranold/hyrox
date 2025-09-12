-- name: CreateStravaAccessToken :one
INSERT INTO strava_tokens (id, strava_user_id, access_token, token_type, expires_at, refresh_token, created_at, updated_at)
VALUES (
    gen_random_uuid(),
    $1,$2,$3,$4,$5,NOW(),NOW()
)
RETURNING  *;

-- name: GetStravaAccessTokenByUserID :one
SELECT * FROM strava_tokens
WHERE strava_user_id = $1;

-- name: GetStravaAccessTokensByApplicationUserID :one
SELECT * FROM strava_tokens st
JOIN strava_user su ON su.id = st.strava_user_id
WHERE su.user_id = $1;

-- name: UpdateStravaAccessToken :exec
UPDATE strava_tokens
SET access_token = $2,
    token_type = $3,
    expires_at = $4,
    refresh_token = $5,
    updated_at = NOW()
Where strava_user_id=$1;

-- name: DeleteAllStravaAccessTokens :exec
DELETE FROM strava_tokens;