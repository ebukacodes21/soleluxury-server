-- name: CreateSession :one
INSERT INTO sessions (
  id, username, user_id,refresh_token, user_agent, client_ip, is_blocked, expired_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;

-- name: Logout :exec
DELETE FROM sessions
WHERE user_id = $1
AND username = $2;

-- name: DeleteSession :exec 
DELETE FROM sessions
WHERE expired_at < NOW();