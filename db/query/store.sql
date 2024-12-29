-- name: CreateStore :one
INSERT INTO stores (
    name
) VALUES (
  $1
)
RETURNING *;

-- name: GetStore :one
SELECT * FROM stores
WHERE id = $1
LIMIT 1;

-- name: GetFirstStore :one
SELECT * FROM stores
ORDER BY created_at ASC 
LIMIT 1;

-- name: GetStores :many
SELECT * FROM stores
ORDER BY id;