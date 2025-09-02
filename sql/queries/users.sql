-- name: CreateUser :one
INSERT INTO users (id, username, email, created_at, updated_at)
VALUES (
    gen_random_uuid(),
    $1,
    $2,
    NOW(),
    NOW()
)
RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users
join STRAVA_USER on users.id = STRAVA_USER.user_id
WHERE users.email = $1;

-- name: GetUserByID :one
SELECT * FROM users
join STRAVA_USER on users.id = STRAVA_USER.user_id
WHERE users.id = $1;
