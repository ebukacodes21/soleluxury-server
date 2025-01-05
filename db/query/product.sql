-- name: CreateProduct :one
INSERT INTO products (
    name, price, is_featured, is_archived, description, images
)
VALUES (
    $1, $2, $3, $4, $5, $6
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

-- name: GetProducts :many
SELECT
    p.id AS product_id,
    p.name AS product_name,
    p.description AS product_description,
    p.price AS product_price,
    p.is_featured AS is_featured,
    p.is_archived AS is_archived,
    p.images AS product_images,
    p.created_at AS product_created_at,
    c.id AS category_id,
    c.name AS category_name,
    cl.id AS color_id,
    cl.value AS color_value,
    sz.id AS size_id,
    sz.value AS size_value
FROM 
    products p
JOIN 
    product_stores ps ON p.id = ps.product_id
LEFT JOIN 
    product_categories pc ON p.id = pc.product_id
LEFT JOIN 
    categories c ON pc.category_id = c.id
LEFT JOIN 
    product_colors pcl ON p.id = pcl.product_id   
LEFT JOIN 
    colors cl ON pcl.color_id = cl.id             
LEFT JOIN 
    product_sizes psz ON p.id = psz.product_id
LEFT JOIN 
    sizes sz ON psz.size_id = sz.id
WHERE 
    ps.store_id = $1
    AND p.is_archived = false;

-- name: GetProduct :one
SELECT
    p.id AS product_id,
    p.name AS product_name,
    p.description AS product_description,
    p.price AS product_price,
    p.is_featured AS is_featured,
    p.is_archived AS is_archived,
    p.images AS product_images,
    p.created_at AS product_created_at,
    c.id AS category_id,
    c.name AS category_name,
    cl.id AS color_id,
    cl.value AS color_value,
    sz.id AS size_id,
    sz.value AS size_value
FROM 
    products p
JOIN 
    product_stores ps ON p.id = ps.product_id
LEFT JOIN 
    product_categories pc ON p.id = pc.product_id
LEFT JOIN 
    categories c ON pc.category_id = c.id
LEFT JOIN 
    product_colors pcl ON p.id = pcl.product_id   
LEFT JOIN 
    colors cl ON pcl.color_id = cl.id             
LEFT JOIN 
    product_sizes psz ON p.id = psz.product_id
LEFT JOIN 
    sizes sz ON psz.size_id = sz.id
WHERE ps.store_id = $1
    AND ps.product_id = $2
LIMIT 1; 

-- name: UpdateProduct :exec
UPDATE products
SET
    name = COALESCE(sqlc.narg(name), name),
    price = COALESCE(sqlc.narg(price), price),
    is_featured = COALESCE(sqlc.narg(is_featured), is_featured),
    is_archived = COALESCE(sqlc.narg(is_archived), is_archived),
    description = COALESCE(sqlc.narg(description), description),
    images = COALESCE(sqlc.narg(images), images)
WHERE 
  id = sqlc.arg(id);

-- name: UpdateProductColor :exec
UPDATE product_colors
SET 
    color_id = COALESCE(sqlc.narg(color_id), color_id)
WHERE 
  product_id = sqlc.arg(product_id);

-- name: UpdateProductSize :exec
UPDATE product_sizes
SET 
    size_id = COALESCE(sqlc.narg(size_id), size_id)
WHERE 
  product_id = sqlc.arg(product_id);

-- name: UpdateProductCategory :exec
UPDATE product_categories
SET 
    category_id = COALESCE(sqlc.narg(category_id), category_id)
WHERE 
  product_id = sqlc.arg(product_id);


-- name: DeleteProductCategory :exec
DELETE FROM product_categories
WHERE product_id = $1;

-- name: DeleteProductColor :exec
DELETE FROM product_colors
WHERE product_id = $1;

-- name: DeleteProductSize :exec
DELETE FROM product_sizes
WHERE product_id = $1;

-- name: DeleteProductStore :exec
DELETE FROM product_stores
WHERE product_id = $1;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1;
