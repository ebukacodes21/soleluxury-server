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
    -- Aggregated Billboards
    json_agg(
        json_build_object(
            'billboard_id', b.id,
            'billboard_label', b.label,
            'billboard_image_url', b.image_url,
            'billboard_created_at', b.created_at
        )
    ) AS billboards,
    -- Aggregated Categories
    json_agg(
        json_build_object(
            'category_id', c.id,
            'category_store_id', c.store_id,
            'category_billboard_id', c.billboard_id,
            'category_name', c.name,
            'category_created_at', c.created_at,
            'category_updated_at', c.updated_at
        )
    ) AS categories,
    -- Aggregated Sizes
    json_agg(
        json_build_object(
            'size_id', sz.id,
            'size_name', sz.name,
            'size_value', sz.value,
            'size_created_at', sz.created_at,
            'size_updated_at', sz.updated_at
        )
    ) AS sizes,
    -- Aggregated Colors
    json_agg(
        json_build_object(
            'color_id', cl.id,
            'color_name', cl.name,
            'color_value', cl.value,
            'color_created_at', cl.created_at,
            'color_updated_at', cl.updated_at
        )
    ) AS colors,
    -- Aggregated Products
    json_agg(
        json_build_object(
            'product_id', p.id,
            'product_name', p.name,
            'product_price', p.price,
            'product_is_featured', p.is_featured,
            'product_is_archived', p.is_archived,
            'product_description', p.description,
            'product_images', p.images,
            'product_created_at', p.created_at,
            'product_updated_at', p.updated_at
        )
    ) AS products,
    -- Aggregated Orders
    json_agg(
        json_build_object(
            'order_id', o.id,
            'order_items', o.items,
            'order_is_paid', o.is_paid,
            'order_phone', o.phone,
            'order_address', o.address,
            'order_created_at', o.created_at,
            'order_updated_at', o.updated_at
        )
    ) AS orders
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
    products p ON s.id = p.store_id
LEFT JOIN 
    orders o ON s.id = o.store_id
WHERE 
    s.id = $1
GROUP BY 
    s.id, s.name, s.created_at;


-- name: GetFirstStore :one
SELECT 
    s.id AS store_id,
    s.name AS store_name,
    s.created_at AS store_created_at,
    -- Aggregated Billboards
    json_agg(
        json_build_object(
            'billboard_id', b.id,
            'billboard_label', b.label,
            'billboard_image_url', b.image_url,
            'billboard_created_at', b.created_at
        )
    ) AS billboards,
    -- Aggregated Categories
    json_agg(
        json_build_object(
            'category_id', c.id,
            'category_store_id', c.store_id,
            'category_billboard_id', c.billboard_id,
            'category_name', c.name,
            'category_created_at', c.created_at,
            'category_updated_at', c.updated_at
        )
    ) AS categories,
    -- Aggregated Sizes
    json_agg(
        json_build_object(
            'size_id', sz.id,
            'size_name', sz.name,
            'size_value', sz.value,
            'size_created_at', sz.created_at,
            'size_updated_at', sz.updated_at
        )
    ) AS sizes,
    -- Aggregated Colors
    json_agg(
        json_build_object(
            'color_id', cl.id,
            'color_name', cl.name,
            'color_value', cl.value,
            'color_created_at', cl.created_at,
            'color_updated_at', cl.updated_at
        )
    ) AS colors,
    -- Aggregated Products
    json_agg(
        json_build_object(
            'product_id', p.id,
            'product_name', p.name,
            'product_price', p.price,
            'product_is_featured', p.is_featured,
            'product_is_archived', p.is_archived,
            'product_description', p.description,
            'product_images', p.images,
            'product_created_at', p.created_at,
            'product_updated_at', p.updated_at
        )
    ) AS products,
    -- Aggregated Orders
    json_agg(
        json_build_object(
            'order_id', o.id,
            'order_items', o.items,
            'order_is_paid', o.is_paid,
            'order_phone', o.phone,
            'order_address', o.address,
            'order_created_at', o.created_at,
            'order_updated_at', o.updated_at
        )
    ) AS orders
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
    products p ON s.id = p.store_id
LEFT JOIN 
    orders o ON s.id = o.store_id
GROUP BY 
    s.id, s.name, s.created_at
ORDER BY 
    s.created_at ASC
LIMIT 1;
 

-- name: GetStores :many
WITH store_data AS (
    SELECT DISTINCT ON (s.id) 
        s.id AS store_id,
        s.name AS store_name,
        s.created_at AS store_created_at
    FROM stores s
    ORDER BY s.id -- Or use ORDER BY s.created_at if you need sorting
)
SELECT 
    sd.store_id,
    sd.store_name,
    sd.store_created_at,
    -- Aggregated Billboards
    (SELECT json_agg(
            json_build_object(
                'billboard_id', b.id,
                'billboard_label', b.label,
                'billboard_image_url', b.image_url,
                'billboard_created_at', b.created_at
            )
        )
     FROM billboards b
     WHERE b.store_id = sd.store_id) AS billboards,
    -- Aggregated Categories
    (SELECT json_agg(
            json_build_object(
                'category_id', c.id,
                'category_store_id', c.store_id,
                'category_billboard_id', c.billboard_id,
                'category_name', c.name,
                'category_created_at', c.created_at,
                'category_updated_at', c.updated_at
            )
        )
     FROM categories c
     WHERE c.store_id = sd.store_id) AS categories,
    -- Aggregated Sizes
    (SELECT json_agg(
            json_build_object(
                'size_id', sz.id,
                'size_name', sz.name,
                'size_value', sz.value,
                'size_created_at', sz.created_at,
                'size_updated_at', sz.updated_at
            )
        )
     FROM sizes sz
     WHERE sz.store_id = sd.store_id) AS sizes,
    -- Aggregated Colors
    (SELECT json_agg(
            json_build_object(
                'color_id', cl.id,
                'color_name', cl.name,
                'color_value', cl.value,
                'color_created_at', cl.created_at,
                'color_updated_at', cl.updated_at
            )
        )
     FROM colors cl
     WHERE cl.store_id = sd.store_id) AS colors,
    -- Aggregated Products
    (SELECT json_agg(
            json_build_object(
                'product_id', p.id,
                'product_name', p.name,
                'product_price', p.price,
                'product_is_featured', p.is_featured,
                'product_is_archived', p.is_archived,
                'product_description', p.description,
                'product_images', p.images,
                'product_created_at', p.created_at,
                'product_updated_at', p.updated_at
            )
        )
     FROM products p
     WHERE p.store_id = sd.store_id) AS products,
    -- Aggregated Orders
    (SELECT json_agg(
            json_build_object(
                'order_id', o.id,
                'order_items', o.items,
                'order_is_paid', o.is_paid,
                'order_phone', o.phone,
                'order_address', o.address,
                'order_created_at', o.created_at,
                'order_updated_at', o.updated_at
            )
        )
     FROM orders o
     WHERE o.store_id = sd.store_id) AS orders
FROM store_data sd
LIMIT $1;


-- name: UpdateStore :exec
UPDATE stores
SET
  name = COALESCE(sqlc.narg(name), name)
WHERE 
  id = sqlc.arg(id);

-- name: DeleteStore :exec
DELETE FROM stores
WHERE id = $1;