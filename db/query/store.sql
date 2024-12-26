-- name: CreateStore :one
INSERT INTO stores (
    name
) VALUES (
  $1
)
RETURNING *;