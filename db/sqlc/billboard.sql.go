// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: billboard.sql

package db

import (
	"context"
)

const createBillboard = `-- name: CreateBillboard :one
INSERT INTO billboards (
  store_id, label, image_url
) VALUES (
  $1, $2, $3
)
RETURNING id, store_id, label, image_url, created_at
`

type CreateBillboardParams struct {
	StoreID  int64  `db:"store_id" json:"store_id"`
	Label    string `db:"label" json:"label"`
	ImageUrl string `db:"image_url" json:"image_url"`
}

func (q *Queries) CreateBillboard(ctx context.Context, arg CreateBillboardParams) (Billboard, error) {
	row := q.db.QueryRowContext(ctx, createBillboard, arg.StoreID, arg.Label, arg.ImageUrl)
	var i Billboard
	err := row.Scan(
		&i.ID,
		&i.StoreID,
		&i.Label,
		&i.ImageUrl,
		&i.CreatedAt,
	)
	return i, err
}

const getBillboard = `-- name: GetBillboard :one
SELECT id, store_id, label, image_url, created_at FROM billboards
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetBillboard(ctx context.Context, id int64) (Billboard, error) {
	row := q.db.QueryRowContext(ctx, getBillboard, id)
	var i Billboard
	err := row.Scan(
		&i.ID,
		&i.StoreID,
		&i.Label,
		&i.ImageUrl,
		&i.CreatedAt,
	)
	return i, err
}
