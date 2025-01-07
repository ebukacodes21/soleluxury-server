-- name: CreateCategory :one
INSERT INTO categories (
  store_id, billboard_id, name
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetCategory :one
SELECT
    c.id AS category_id,
    c.store_id AS category_store_id,
    c.billboard_id AS category_billboard_id,
    c.name AS category_name,
    c.created_at AS category_created_at,
    c.updated_at AS category_updated_at,
    -- Associated Billboard Details
    json_build_object(
        'billboard_id', b.id,
        'billboard_label', b.label,
        'billboard_image_url', b.image_url,
        'billboard_created_at', b.created_at
    ) AS billboard
FROM
    categories c
LEFT JOIN
    billboards b ON c.billboard_id = b.id
WHERE
    c.id = $1
LIMIT 1;

-- name: GetCategories :many
SELECT
    c.id AS category_id,
    c.store_id AS category_store_id,
    c.billboard_id AS category_billboard_id,
    c.name AS category_name,
    c.created_at AS category_created_at,
    c.updated_at AS category_updated_at,
    -- Associated Billboard Details
    json_agg(
        json_build_object(
            'billboard_id', b.id,
            'billboard_label', b.label,
            'billboard_image_url', b.image_url,
            'billboard_created_at', b.created_at
        )
    ) AS billboards
FROM
    categories c
LEFT JOIN
    billboards b ON c.billboard_id = b.id
WHERE
    c.store_id = $1
GROUP BY
    c.id
ORDER BY
    c.id;

-- name: UpdateCategory :exec
UPDATE categories
SET
  name = COALESCE(sqlc.narg(name), name)
WHERE 
  id = sqlc.arg(id)
  AND store_id = sqlc.arg(store_id);

-- name: DeleteCategory :exec
DELETE FROM categories
WHERE id = $1;