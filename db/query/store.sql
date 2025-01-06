-- name: CreateStore :one
INSERT INTO stores (
    name
) VALUES (
  $1
) RETURNING *;

-- name: GetStore :one
SELECT
    s.id AS store_id,
    s.name AS store_name,
    s.created_at AS store_created_at,
    -- Billboards
    b.id AS billboard_id,
    b.label AS billboard_label,
    b.image_url AS billboard_image_url,
    b.created_at AS billboard_created_at,
    -- Categories
    c.id AS category_id,
    c.store_name AS category_store_name,
    c.billboard_label AS category_billboard_label,
    c.name AS category_name,
    c.created_at AS category_created_at,
    c.updated_at AS category_updated_at,
    -- Sizes
    sz.id AS size_id,
    sz.store_name AS size_store_name,
    sz.name AS size_name,
    sz.value AS size_value,
    sz.created_at AS size_created_at,
    sz.updated_at AS size_updated_at,
    -- Colors
    cl.id AS color_id,
    cl.store_name AS color_store_name,
    cl.name AS color_name,
    cl.value AS color_value,
    cl.created_at AS color_created_at,
    cl.updated_at AS color_updated_at,
    -- Products (using product_stores junction table)
    p.id AS product_id,
    p.name AS product_name,
    p.price AS product_price,
    p.is_featured AS product_is_featured,
    p.is_archived AS product_is_archived,
    p.description AS product_description,
    p.images AS product_images,
    p.created_at AS product_created_at,
    p.updated_at AS product_updated_at,
    -- Orders
    o.id AS order_id,
    o.items AS order_items,
    o.is_paid AS order_is_paid,
    o.phone AS order_phone,
    o.address AS order_address,
    o.created_at AS order_created_at,
    o.updated_at AS order_updated_at
FROM
    stores s
LEFT JOIN
    billboards b ON s.id = b.store_id
LEFT JOIN
    categories c ON s.id = c.store_id
LEFT JOIN
    sizes sz ON s.id = sz.store_id
LEFT JOIN
    colors cl ON s.id = cl.store_id
LEFT JOIN
    orders o ON s.id = o.store_id
LEFT JOIN
    product_stores ps ON s.id = ps.store_id
LEFT JOIN
    products p ON ps.product_id = p.id
WHERE
    s.id = $1; 

-- name: GetFirstStore :one
SELECT
    s.id AS store_id,
    s.name AS store_name,
    s.created_at AS store_created_at,
    -- Billboards
    b.id AS billboard_id,
    b.label AS billboard_label,
    b.image_url AS billboard_image_url,
    b.created_at AS billboard_created_at,
    -- Categories
    c.id AS category_id,
    c.store_name AS category_store_name,
    c.billboard_label AS category_billboard_label,
    c.name AS category_name,
    c.created_at AS category_created_at,
    c.updated_at AS category_updated_at,
    -- Sizes
    sz.id AS size_id,
    sz.store_name AS size_store_name,
    sz.name AS size_name,
    sz.value AS size_value,
    sz.created_at AS size_created_at,
    sz.updated_at AS size_updated_at,
    -- Colors
    cl.id AS color_id,
    cl.store_name AS color_store_name,
    cl.name AS color_name,
    cl.value AS color_value,
    cl.created_at AS color_created_at,
    cl.updated_at AS color_updated_at,
    -- Products (using product_stores junction table)
    p.id AS product_id,
    p.name AS product_name,
    p.price AS product_price,
    p.is_featured AS product_is_featured,
    p.is_archived AS product_is_archived,
    p.description AS product_description,
    p.images AS product_images,
    p.created_at AS product_created_at,
    p.updated_at AS product_updated_at,
    -- Orders
    o.id AS order_id,
    o.items AS order_items,
    o.is_paid AS order_is_paid,
    o.phone AS order_phone,
    o.address AS order_address,
    o.created_at AS order_created_at,
    o.updated_at AS order_updated_at
FROM
    stores s
LEFT JOIN
    billboards b ON s.id = b.store_id
LEFT JOIN
    categories c ON s.id = c.store_id
LEFT JOIN
    sizes sz ON s.id = sz.store_id
LEFT JOIN
    colors cl ON s.id = cl.store_id
LEFT JOIN
    orders o ON s.id = o.store_id
LEFT JOIN
    product_stores ps ON s.id = ps.store_id
LEFT JOIN
    products p ON ps.product_id = p.id
ORDER BY
    s.created_at ASC
LIMIT 1;

-- name: GetStores :many
SELECT
    s.id AS store_id,
    s.name AS store_name,
    s.created_at AS store_created_at,
    -- Billboards
    b.id AS billboard_id,
    b.label AS billboard_label,
    b.image_url AS billboard_image_url,
    b.created_at AS billboard_created_at,
    -- Categories
    c.id AS category_id,
    c.store_name AS category_store_name,
    c.billboard_label AS category_billboard_label,
    c.name AS category_name,
    c.created_at AS category_created_at,
    c.updated_at AS category_updated_at,
    -- Sizes
    sz.id AS size_id,
    sz.store_name AS size_store_name,
    sz.name AS size_name,
    sz.value AS size_value,
    sz.created_at AS size_created_at,
    sz.updated_at AS size_updated_at,
    -- Colors
    cl.id AS color_id,
    cl.store_name AS color_store_name,
    cl.name AS color_name,
    cl.value AS color_value,
    cl.created_at AS color_created_at,
    cl.updated_at AS color_updated_at,
    -- Products (using product_stores junction table)
    p.id AS product_id,
    p.name AS product_name,
    p.price AS product_price,
    p.is_featured AS product_is_featured,
    p.is_archived AS product_is_archived,
    p.description AS product_description,
    p.images AS product_images,
    p.created_at AS product_created_at,
    p.updated_at AS product_updated_at,
    -- Orders
    o.id AS order_id,
    o.items AS order_items,
    o.is_paid AS order_is_paid,
    o.phone AS order_phone,
    o.address AS order_address,
    o.created_at AS order_created_at,
    o.updated_at AS order_updated_at
FROM
    stores s
LEFT JOIN
    billboards b ON s.id = b.store_id
LEFT JOIN
    categories c ON s.id = c.store_id
LEFT JOIN
    sizes sz ON s.id = sz.store_id
LEFT JOIN
    colors cl ON s.id = cl.store_id
LEFT JOIN
    orders o ON s.id = o.store_id
LEFT JOIN
    product_stores ps ON s.id = ps.store_id
LEFT JOIN
    products p ON ps.product_id = p.id
ORDER BY
    s.created_at DESC
LIMIT $1;  -- Here $1 is the limit passed (number of stores to fetch)   

-- name: UpdateStore :exec
UPDATE stores
SET
  name = COALESCE(sqlc.narg(name), name)
WHERE 
  id = sqlc.arg(id);

-- name: DeleteStore :exec
DELETE FROM stores
WHERE id = $1;