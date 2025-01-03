-- name: CreateSize :one
INSERT INTO sizes (
  store_id, store_name, name, value
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetSize :one
SELECT * FROM sizes
WHERE id = $1
LIMIT 1;

-- name: GetSizes :many
SELECT * FROM sizes
WHERE store_id = $1
ORDER BY id;

-- name: UpdateSize :exec
UPDATE sizes
SET
  name = COALESCE(sqlc.narg(name), name),
  value = COALESCE(sqlc.narg(value), value)
WHERE 
  id = sqlc.arg(id)
  AND store_id = sqlc.arg(store_id);

-- name: DeleteSize :exec
DELETE FROM sizes
WHERE id = $1;