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

-- name: GetBillboards :many
SELECT * FROM billboards
WHERE store_id = $1
ORDER BY id;

-- name: UpdateBillboard :exec
UPDATE billboards
SET
  label = COALESCE(sqlc.narg(label), label),
  image_url = COALESCE(sqlc.narg(image_url), image_url)
WHERE 
  id = sqlc.arg(id)
  AND store_id = sqlc.arg(store_id);

-- name: DeleteBillboard :exec
DELETE FROM billboards
WHERE id = $1;