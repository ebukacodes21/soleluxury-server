-- name: CreateOrder :one
INSERT INTO orders (
    store_id
)
VALUES (
    $1
)
RETURNING *;

-- name: GetOrders :many
SELECT * FROM orders
WHERE store_id = $1
ORDER BY id;