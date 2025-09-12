-- name: CreateUser :one
INSERT INTO users (id, username, email, hashed_password, created_at, updated_at)
VALUES (
    gen_random_uuid(),
    $1,
    $2,
    $3,
    NOW(),
    NOW()
)
RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: GetUserByID :one
SELECT * FROM users
join STRAVA_USER on users.id = STRAVA_USER.user_id
WHERE users.id = $1;

-- name: UpdatePassword :exec
UPDATE users
Set hashed_password = $2,
    updated_at = NOW()
Where id=$1;

-- name: DeleteAllUsers :exec
DELETE FROM users;
