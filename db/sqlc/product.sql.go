// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: product.sql

package db

import (
	"context"
	"database/sql"
	"encoding/json"
)

const createImage = `-- name: CreateImage :one
INSERT INTO images (
    product_id, urls
)
VALUES (
    $1, $2
)
RETURNING id, product_id, urls, created_at
`

type CreateImageParams struct {
	ProductID int64           `db:"product_id" json:"product_id"`
	Urls      json.RawMessage `db:"urls" json:"urls"`
}

func (q *Queries) CreateImage(ctx context.Context, arg CreateImageParams) (Image, error) {
	row := q.db.QueryRowContext(ctx, createImage, arg.ProductID, arg.Urls)
	var i Image
	err := row.Scan(
		&i.ID,
		&i.ProductID,
		&i.Urls,
		&i.CreatedAt,
	)
	return i, err
}

const createProduct = `-- name: CreateProduct :one
INSERT INTO products (
    name, price, is_featured, is_archived, description
)
VALUES (
    $1, $2, $3, $4, $5
)
RETURNING id, name, price, is_featured, is_archived, description, created_at, updated_at
`

type CreateProductParams struct {
	Name        string         `db:"name" json:"name"`
	Price       float64        `db:"price" json:"price"`
	IsFeatured  bool           `db:"is_featured" json:"is_featured"`
	IsArchived  bool           `db:"is_archived" json:"is_archived"`
	Description sql.NullString `db:"description" json:"description"`
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error) {
	row := q.db.QueryRowContext(ctx, createProduct,
		arg.Name,
		arg.Price,
		arg.IsFeatured,
		arg.IsArchived,
		arg.Description,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Price,
		&i.IsFeatured,
		&i.IsArchived,
		&i.Description,
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

const createProductStore = `-- name: CreateProductStore :one
INSERT INTO product_stores (
    product_id, store_id
)
VALUES (
    $1, $2
)
RETURNING product_id, store_id
`

type CreateProductStoreParams struct {
	ProductID int64 `db:"product_id" json:"product_id"`
	StoreID   int64 `db:"store_id" json:"store_id"`
}

func (q *Queries) CreateProductStore(ctx context.Context, arg CreateProductStoreParams) (ProductStore, error) {
	row := q.db.QueryRowContext(ctx, createProductStore, arg.ProductID, arg.StoreID)
	var i ProductStore
	err := row.Scan(&i.ProductID, &i.StoreID)
	return i, err
}