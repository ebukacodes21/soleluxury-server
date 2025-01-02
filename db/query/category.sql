-- name: CreateCategory :one
INSERT INTO categories (
  store_id, billboard_id, name
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetCategory :one
SELECT * FROM categories
WHERE id = $1
LIMIT 1;