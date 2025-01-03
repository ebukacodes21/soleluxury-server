-- name: CreateUser :one
INSERT INTO users (
    username, email, password, verification_code
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users 
WHERE email = $1
LIMIT 1;