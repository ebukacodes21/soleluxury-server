// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: product.sql

package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/sqlc-dev/pqtype"
)

const createProduct = `-- name: CreateProduct :one
INSERT INTO products (
    store_id, name, price, is_featured, is_archived, description, images
)
VALUES (
    $1, $2, $3, $4, $5, $6, $7
)
RETURNING id, store_id, name, price, is_featured, is_archived, description, images, created_at, updated_at
`

type CreateProductParams struct {
	StoreID     int64           `db:"store_id" json:"store_id"`
	Name        string          `db:"name" json:"name"`
	Price       float64         `db:"price" json:"price"`
	IsFeatured  bool            `db:"is_featured" json:"is_featured"`
	IsArchived  bool            `db:"is_archived" json:"is_archived"`
	Description string          `db:"description" json:"description"`
	Images      json.RawMessage `db:"images" json:"images"`
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error) {
	row := q.db.QueryRowContext(ctx, createProduct,
		arg.StoreID,
		arg.Name,
		arg.Price,
		arg.IsFeatured,
		arg.IsArchived,
		arg.Description,
		arg.Images,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.StoreID,
		&i.Name,
		&i.Price,
		&i.IsFeatured,
		&i.IsArchived,
		&i.Description,
		&i.Images,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createProductCategory = `-- name: CreateProductCategory :one
INSERT INTO product_categories (
    product_id, category_id
)
VALUES (
    $1, $2
)
RETURNING product_id, category_id
`

type CreateProductCategoryParams struct {
	ProductID  int64 `db:"product_id" json:"product_id"`
	CategoryID int64 `db:"category_id" json:"category_id"`
}

func (q *Queries) CreateProductCategory(ctx context.Context, arg CreateProductCategoryParams) (ProductCategory, error) {
	row := q.db.QueryRowContext(ctx, createProductCategory, arg.ProductID, arg.CategoryID)
	var i ProductCategory
	err := row.Scan(&i.ProductID, &i.CategoryID)
	return i, err
}

const createProductColor = `-- name: CreateProductColor :one
INSERT INTO product_colors (
    product_id, color_id
)
VALUES (
    $1, $2
)
RETURNING product_id, color_id
`

type CreateProductColorParams struct {
	ProductID int64 `db:"product_id" json:"product_id"`
	ColorID   int64 `db:"color_id" json:"color_id"`
}

func (q *Queries) CreateProductColor(ctx context.Context, arg CreateProductColorParams) (ProductColor, error) {
	row := q.db.QueryRowContext(ctx, createProductColor, arg.ProductID, arg.ColorID)
	var i ProductColor
	err := row.Scan(&i.ProductID, &i.ColorID)
	return i, err
}

const createProductSize = `-- name: CreateProductSize :one
INSERT INTO product_sizes (
    product_id, size_id
)
VALUES (
    $1, $2
)
RETURNING product_id, size_id
`

type CreateProductSizeParams struct {
	ProductID int64 `db:"product_id" json:"product_id"`
	SizeID    int64 `db:"size_id" json:"size_id"`
}

func (q *Queries) CreateProductSize(ctx context.Context, arg CreateProductSizeParams) (ProductSize, error) {
	row := q.db.QueryRowContext(ctx, createProductSize, arg.ProductID, arg.SizeID)
	var i ProductSize
	err := row.Scan(&i.ProductID, &i.SizeID)
	return i, err
}

const deleteProduct = `-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1
`

func (q *Queries) DeleteProduct(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteProduct, id)
	return err
}

const deleteProductCategory = `-- name: DeleteProductCategory :exec
DELETE FROM product_categories
WHERE product_id = $1
`

func (q *Queries) DeleteProductCategory(ctx context.Context, productID int64) error {
	_, err := q.db.ExecContext(ctx, deleteProductCategory, productID)
	return err
}

const deleteProductColor = `-- name: DeleteProductColor :exec
DELETE FROM product_colors
WHERE product_id = $1
`

func (q *Queries) DeleteProductColor(ctx context.Context, productID int64) error {
	_, err := q.db.ExecContext(ctx, deleteProductColor, productID)
	return err
}

const deleteProductSize = `-- name: DeleteProductSize :exec
DELETE FROM product_sizes
WHERE product_id = $1
`

func (q *Queries) DeleteProductSize(ctx context.Context, productID int64) error {
	_, err := q.db.ExecContext(ctx, deleteProductSize, productID)
	return err
}

const getProduct = `-- name: GetProduct :one
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
WHERE p.store_id = $1
    AND p.id = $2
LIMIT 1
`

type GetProductParams struct {
	StoreID int64 `db:"store_id" json:"store_id"`
	ID      int64 `db:"id" json:"id"`
}

type GetProductRow struct {
	ProductID          int64           `db:"product_id" json:"product_id"`
	ProductName        string          `db:"product_name" json:"product_name"`
	ProductDescription string          `db:"product_description" json:"product_description"`
	ProductPrice       float64         `db:"product_price" json:"product_price"`
	IsFeatured         bool            `db:"is_featured" json:"is_featured"`
	IsArchived         bool            `db:"is_archived" json:"is_archived"`
	ProductImages      json.RawMessage `db:"product_images" json:"product_images"`
	ProductCreatedAt   time.Time       `db:"product_created_at" json:"product_created_at"`
	CategoryID         sql.NullInt64   `db:"category_id" json:"category_id"`
	CategoryName       sql.NullString  `db:"category_name" json:"category_name"`
	ColorID            sql.NullInt64   `db:"color_id" json:"color_id"`
	ColorValue         sql.NullString  `db:"color_value" json:"color_value"`
	SizeID             sql.NullInt64   `db:"size_id" json:"size_id"`
	SizeValue          sql.NullString  `db:"size_value" json:"size_value"`
}

func (q *Queries) GetProduct(ctx context.Context, arg GetProductParams) (GetProductRow, error) {
	row := q.db.QueryRowContext(ctx, getProduct, arg.StoreID, arg.ID)
	var i GetProductRow
	err := row.Scan(
		&i.ProductID,
		&i.ProductName,
		&i.ProductDescription,
		&i.ProductPrice,
		&i.IsFeatured,
		&i.IsArchived,
		&i.ProductImages,
		&i.ProductCreatedAt,
		&i.CategoryID,
		&i.CategoryName,
		&i.ColorID,
		&i.ColorValue,
		&i.SizeID,
		&i.SizeValue,
	)
	return i, err
}

const getProducts = `-- name: GetProducts :many
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
    p.store_id = $1
    AND p.is_archived = false
`

type GetProductsRow struct {
	ProductID          int64           `db:"product_id" json:"product_id"`
	ProductName        string          `db:"product_name" json:"product_name"`
	ProductDescription string          `db:"product_description" json:"product_description"`
	ProductPrice       float64         `db:"product_price" json:"product_price"`
	IsFeatured         bool            `db:"is_featured" json:"is_featured"`
	IsArchived         bool            `db:"is_archived" json:"is_archived"`
	ProductImages      json.RawMessage `db:"product_images" json:"product_images"`
	ProductCreatedAt   time.Time       `db:"product_created_at" json:"product_created_at"`
	CategoryID         sql.NullInt64   `db:"category_id" json:"category_id"`
	CategoryName       sql.NullString  `db:"category_name" json:"category_name"`
	ColorID            sql.NullInt64   `db:"color_id" json:"color_id"`
	ColorValue         sql.NullString  `db:"color_value" json:"color_value"`
	SizeID             sql.NullInt64   `db:"size_id" json:"size_id"`
	SizeValue          sql.NullString  `db:"size_value" json:"size_value"`
}

func (q *Queries) GetProducts(ctx context.Context, storeID int64) ([]GetProductsRow, error) {
	rows, err := q.db.QueryContext(ctx, getProducts, storeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetProductsRow{}
	for rows.Next() {
		var i GetProductsRow
		if err := rows.Scan(
			&i.ProductID,
			&i.ProductName,
			&i.ProductDescription,
			&i.ProductPrice,
			&i.IsFeatured,
			&i.IsArchived,
			&i.ProductImages,
			&i.ProductCreatedAt,
			&i.CategoryID,
			&i.CategoryName,
			&i.ColorID,
			&i.ColorValue,
			&i.SizeID,
			&i.SizeValue,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateProduct = `-- name: UpdateProduct :exec
UPDATE products
SET
    name = COALESCE($1, name),
    price = COALESCE($2, price),
    is_featured = COALESCE($3, is_featured),
    is_archived = COALESCE($4, is_archived),
    description = COALESCE($5, description),
    images = COALESCE($6, images)
WHERE 
  id = $7
`

type UpdateProductParams struct {
	Name        sql.NullString        `db:"name" json:"name"`
	Price       sql.NullFloat64       `db:"price" json:"price"`
	IsFeatured  sql.NullBool          `db:"is_featured" json:"is_featured"`
	IsArchived  sql.NullBool          `db:"is_archived" json:"is_archived"`
	Description sql.NullString        `db:"description" json:"description"`
	Images      pqtype.NullRawMessage `db:"images" json:"images"`
	ID          int64                 `db:"id" json:"id"`
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) error {
	_, err := q.db.ExecContext(ctx, updateProduct,
		arg.Name,
		arg.Price,
		arg.IsFeatured,
		arg.IsArchived,
		arg.Description,
		arg.Images,
		arg.ID,
	)
	return err
}

const updateProductCategory = `-- name: UpdateProductCategory :exec
UPDATE product_categories
SET 
    category_id = COALESCE($1, category_id)
WHERE 
  product_id = $2
`

type UpdateProductCategoryParams struct {
	CategoryID sql.NullInt64 `db:"category_id" json:"category_id"`
	ProductID  int64         `db:"product_id" json:"product_id"`
}

func (q *Queries) UpdateProductCategory(ctx context.Context, arg UpdateProductCategoryParams) error {
	_, err := q.db.ExecContext(ctx, updateProductCategory, arg.CategoryID, arg.ProductID)
	return err
}

const updateProductColor = `-- name: UpdateProductColor :exec
UPDATE product_colors
SET 
    color_id = COALESCE($1, color_id)
WHERE 
  product_id = $2
`

type UpdateProductColorParams struct {
	ColorID   sql.NullInt64 `db:"color_id" json:"color_id"`
	ProductID int64         `db:"product_id" json:"product_id"`
}

func (q *Queries) UpdateProductColor(ctx context.Context, arg UpdateProductColorParams) error {
	_, err := q.db.ExecContext(ctx, updateProductColor, arg.ColorID, arg.ProductID)
	return err
}

const updateProductSize = `-- name: UpdateProductSize :exec
UPDATE product_sizes
SET 
    size_id = COALESCE($1, size_id)
WHERE 
  product_id = $2
`

type UpdateProductSizeParams struct {
	SizeID    sql.NullInt64 `db:"size_id" json:"size_id"`
	ProductID int64         `db:"product_id" json:"product_id"`
}

func (q *Queries) UpdateProductSize(ctx context.Context, arg UpdateProductSizeParams) error {
	_, err := q.db.ExecContext(ctx, updateProductSize, arg.SizeID, arg.ProductID)
	return err
}
