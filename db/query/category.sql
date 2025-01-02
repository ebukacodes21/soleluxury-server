-- name: CreateCategory :one
INSERT INTO categories (
  store_id, billboard_id, store_name, billboard_label, name
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetCategory :one
SELECT * FROM categories
WHERE id = $1
LIMIT 1;

-- name: GetCategories :many
SELECT * FROM categories
WHERE store_id = $1
ORDER BY id;

-- name: UpdateCategory :exec
UPDATE categories
SET
  name = COALESCE(sqlc.narg(name), name),
  billboard_label = COALESCE(sqlc.narg(billboard_label), billboard_label)
WHERE 
  id = sqlc.arg(id)
  AND store_id = sqlc.arg(store_id);

-- name: DeleteCategory :exec
DELETE FROM categories
WHERE id = $1;