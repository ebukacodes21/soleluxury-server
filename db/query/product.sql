-- name: CreateProduct :one
INSERT INTO products (
    name, price, is_featured, is_archived, description
)
VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: CreateImage :one
INSERT INTO images (
    product_id, urls
)
VALUES (
    $1, $2
)
RETURNING *;

-- name: CreateProductColor :one
INSERT INTO product_colors (
    product_id, color_id
)
VALUES (
    $1, $2
)
RETURNING *;

-- name: CreateProductSize :one
INSERT INTO product_sizes (
    product_id, size_id
)
VALUES (
    $1, $2
)
RETURNING *;

-- name: CreateProductStore :one
INSERT INTO product_stores (
    product_id, store_id
)
VALUES (
    $1, $2
)
RETURNING *;

-- name: CreateProductCategory :one
INSERT INTO product_categories (
    product_id, category_id
)
VALUES (
    $1, $2
)
RETURNING *;
