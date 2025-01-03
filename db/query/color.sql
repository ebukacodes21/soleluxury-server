-- name: CreateColor :one
INSERT INTO colors (
  store_id, store_name, name, value
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetColor :one
SELECT * FROM colors
WHERE id = $1
LIMIT 1;

-- name: GetColors :many
SELECT * FROM colors
WHERE store_id = $1
ORDER BY id;

-- name: UpdateColor :exec
UPDATE colors
SET
  name = COALESCE(sqlc.narg(name), name),
  value = COALESCE(sqlc.narg(value), value)
WHERE 
  id = sqlc.arg(id)
  AND store_id = sqlc.arg(store_id);

-- name: DeleteColor :exec
DELETE FROM colors
WHERE id = $1;