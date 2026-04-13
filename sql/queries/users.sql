-- name: CreateUser :one
INSERT INTO users (
    id, name, created_at, updated_at
) VALUES (
    gen_random_uuid(),
    $1,
    NOW(),
    NOW()
)
RETURNING *;

-- name: GetUserByName :one
SELECT * FROM users
WHERE name = $1;