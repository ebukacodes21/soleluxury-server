-- name: CreateBillboard :one
INSERT INTO billboards (
  store_id, label, image_url
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetBillboard :one
SELECT * FROM billboards
WHERE id = $1
LIMIT 1;